package chat

import (
	"encoding/json"
	"fmt"

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
		userID := ep.Kws.Params("user_id")
		i.mu.RLock()
		i.connections[userID] = ep.Kws.UUID
		i.mu.RUnlock()
		log.Infof("Connected: %s", userID)
	})

	socketio.On("disconnect", func(ep *socketio.EventPayload) {
		userID := ep.Kws.Params("user_id")
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

	socketio.On("join", func(ep *socketio.EventPayload) {
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

		i.addUserToChat(ep.Kws.UUID, chatID)

		log.Info("Join event successful")
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

		ep.Kws.EmitToList(recipients, []byte(message))
	})

	return func(kws *socketio.Websocket) {
		log.Infof("KWS: %+v", kws)
	}
}

func (i *Implementation) addUserToChat(userUUID string, chatID string) {
	i.userToChat[userUUID] = chatID
	i.chatToUsers[chatID] = append(i.chatToUsers[chatID], userUUID)
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
