package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2/log"
)

type MessageObject struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func (i *Implementation) socketErr(ep *socketio.EventPayload, err error) {
	errByte := []byte(err.Error())
	ep.Kws.Emit(errByte)
}

func (i *Implementation) encodeMessage(ep *socketio.EventPayload) (map[string]interface{}, error) {
	message := new(MessageObject)
	if err := json.Unmarshal(ep.Data, message); err != nil {
		return nil, err
	}

	dataMap, ok := message.Data.(map[string]interface{})
	if !ok {
		i.socketErr(ep, nil)
		return nil, fmt.Errorf("failed to parse data")
	}

	return dataMap, nil
}

func (i *Implementation) WS() func(*socketio.Websocket) {
	socketio.On("connect", func(ep *socketio.EventPayload) {
		userID := i.getUserID(ep)
		if userID == "" {
			i.socketErr(ep, fmt.Errorf("user not found"))
			return
		}

		i.mu.RLock()
		i.connections[userID] = ep.Kws.UUID
		i.mu.RUnlock()
		log.Infof("Connected: %s", userID)
	})

	socketio.On("disconnect", func(ep *socketio.EventPayload) {
		userID := i.getUserID(ep)
		if userID == "" {
			i.socketErr(ep, fmt.Errorf("user not found"))
			return
		}

		i.mu.Lock()
		defer i.mu.Unlock()

		delete(i.connections, userID)
		i.removeUserFromChat(ep.Kws.UUID)

		log.Infof("Disconnected: %s", userID)
	})

	socketio.On("message", func(ep *socketio.EventPayload) {
		message := new(MessageObject)
		if err := json.Unmarshal(ep.Data, message); err != nil {
			i.socketErr(ep, err)
		}

		if message.Event != "" {
			ep.Kws.Fire(message.Event, ep.Data)
		}
	})

	// ================ Custom Events ================
	socketio.On("join", func(ep *socketio.EventPayload) { // Подключение к чату обычного пользователя
		userID := i.getUserID(ep)

		if err := i.service.Create(context.Background(), &model.Chat{ID: userID}); err != nil {
			i.socketErr(ep, err)
			return
		}

		i.addUserToChat(ep.Kws.UUID, userID)

		historyMessages, err := i.service.GetMessagesByChatID(context.Background(), userID)
		if err != nil {
			i.socketErr(ep, err)
			return
		}

		jsonData, err := json.Marshal(historyMessages)
		if err != nil {
			i.socketErr(ep, err)
			return
		}

		ep.Kws.EmitTo(ep.Kws.UUID, jsonData)
		log.Info("Join event successful")
	})

	socketio.On("admin-join", func(ep *socketio.EventPayload) {
		dataMap, err := i.encodeMessage(ep)
		if err != nil {
			i.socketErr(ep, err)
			return
		}

		chatID, ok := dataMap["chat_id"].(string)
		if !ok {
			i.socketErr(ep, fmt.Errorf("failed to parse chat_id"))
			return
		}

		existChat, err := i.service.GetByID(context.Background(), chatID)
		if err != nil {
			i.socketErr(ep, err)
			return
		}

		if existChat == nil {
			i.socketErr(ep, fmt.Errorf("chat not found"))
			return
		}

		user, ok := ep.Kws.Locals("user").(response.UserResponse)
		if !ok {
			log.Infof("User not found: %v", ep.Kws.Locals("user"))
			i.socketErr(ep, fmt.Errorf("user not found"))
			return
		}

		if user.Role == "admin" {
			i.addUserToChat(ep.Kws.UUID, chatID)
		} else {
			i.socketErr(ep, fmt.Errorf("user has not permissions"))
			return
		}

		historyMessages, err := i.service.GetMessagesByChatID(context.Background(), chatID)
		if err != nil {
			i.socketErr(ep, err)
			return
		}

		jsonData, err := json.Marshal(historyMessages)
		if err != nil {
			i.socketErr(ep, err)
			return
		}

		ep.Kws.EmitTo(ep.Kws.UUID, jsonData)
	})

	socketio.On("message-event", func(ep *socketio.EventPayload) {
		dataMap, err := i.encodeMessage(ep)
		if err != nil {
			i.socketErr(ep, err)
			return
		}

		chatID, ok := i.userToChat[ep.Kws.UUID]
		if !ok {
			i.socketErr(ep, fmt.Errorf("user not in chat"))
			return
		}

		participants, ok := i.chatToUsers[chatID]
		if !ok {
			i.socketErr(ep, fmt.Errorf("chat not found"))
			return
		}

		var recipients []string
		for _, uuid := range participants {
			if uuid != ep.Kws.UUID {
				recipients = append(recipients, uuid)
			}
		}

		message, ok := dataMap["message"].(string)
		if !ok {
			i.socketErr(ep, fmt.Errorf("failed to parse message"))
			return
		}

		userID := i.getUserID(ep)
		saveMessageData := model.Message{
			SenderID: userID,
			Message:  message,
			ChatID:   chatID,
		}

		if err := i.service.SaveMessage(context.Background(), &saveMessageData); err != nil {
			i.socketErr(ep, err)
			return
		}

		ep.Kws.EmitToList(recipients, []byte(message))
	})

	return func(kws *socketio.Websocket) {
		// log.Infof("KWS: %+v", kws)
	}
}

// ================ Helpers ================
func (i *Implementation) addUserToChat(userUUID string, chatID string) {
	i.userToChat[userUUID] = chatID
	i.chatToUsers[chatID] = append(i.chatToUsers[chatID], userUUID)
}

func (i *Implementation) getUserID(ep *socketio.EventPayload) string {
	user, _ := ep.Kws.Locals("user").(response.UserResponse)
	userID := user.ID

	if userID == "" {
		userID = ep.Kws.Locals("session_id").(string)
	}

	return userID
}

func (i *Implementation) removeUserFromChat(userUUID string) {
	chatID, ok := i.userToChat[userUUID]
	if !ok {
		return // Пользователь не в чате
	}

	users := i.chatToUsers[chatID]
	for index, uuid := range users {
		if uuid == userUUID {
			i.chatToUsers[chatID] = append(users[:index], users[index+1:]...)
			break
		}
	}

	delete(i.userToChat, userUUID)
}
