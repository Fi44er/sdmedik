package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
)

// Message represents a chat message
type Message struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// Chat stores messages and metadata for a support chat
type Chat struct {
	ID           string
	Messages     []Message
	LastActive   time.Time
	Authorized   bool
	MessageMutex sync.Mutex
}

// ChatManager manages all chats (authorized persistent and guest temporary)
type ChatManager struct {
	chats    map[string]*Chat
	mutex    sync.Mutex
	guestTTL time.Duration
}

func NewChatManager(guestTTL time.Duration) *ChatManager {
	cm := &ChatManager{
		chats:    make(map[string]*Chat),
		guestTTL: guestTTL,
	}
	// Start a background goroutine to clean up expired guest chats
	go cm.cleanupExpiredGuests()
	return cm
}

// Get or create chat by ID and auth status
func (cm *ChatManager) GetOrCreateChat(chatID string, authorized bool) *Chat {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	chat, exists := cm.chats[chatID]
	if !exists {
		chat = &Chat{
			ID:         chatID,
			Messages:   []Message{},
			LastActive: time.Now(),
			Authorized: authorized,
		}
		cm.chats[chatID] = chat
	}
	return chat
}

// Add message to chat, updates LastActive time
func (c *Chat) AddMessage(msg Message) {
	c.MessageMutex.Lock()
	defer c.MessageMutex.Unlock()

	c.Messages = append(c.Messages, msg)
	c.LastActive = time.Now()
}

// Cleanup guest chats inactive for more than TTL
func (cm *ChatManager) cleanupExpiredGuests() {
	for {
		time.Sleep(time.Minute)
		cutoff := time.Now().Add(-cm.guestTTL)
		cm.mutex.Lock()
		for id, chat := range cm.chats {
			if !chat.Authorized && chat.LastActive.Before(cutoff) {
				delete(cm.chats, id)
				log.Printf("Deleted guest chat %s due to inactivity\n", id)
			}
		}
		cm.mutex.Unlock()
	}
}

// Simulated user authentication - in real app replace with real auth logic
func authenticateUser(c *fiber.Ctx) (userID, username string, authorized bool) {
	// For demonstration, we check for a query param "user"
	user := c.Query("user", "")
	if user != "" {
		return user, fmt.Sprintf("User-%s", user), true
	}
	return "", "Guest", false
}

func main() {
	app := fiber.New()

	// Initialize socketio server
	server := socketio.New(socketio.Websocket, &socketio.Options{
		PingTimeout: 10 * time.Second,
	})

	chatManager := NewChatManager(30 * time.Minute)

	// On connect handler
	server.OnConnect("/", func(s socketio.Conn) error {
		// Get query params for chatID and user info
		chatID := s.URL().Query().Get("chat_id")
		userID := s.URL().Query().Get("user_id")
		username := s.URL().Query().Get("username")
		isAdmin := s.URL().Query().Get("admin") == "1"

		if chatID == "" {
			return fmt.Errorf("chat_id is required")
		}

		if username == "" {
			username = "Guest"
		}

		s.SetContext(map[string]string{
			"chat_id":  chatID,
			"user_id":  userID,
			"username": username,
			"admin":    fmt.Sprintf("%v", isAdmin),
		})

		s.Join(chatID)

		log.Printf("Connected: user %s joined chat %s (admin: %v)", username, chatID, isAdmin)

		// On connection, send existing messages in chat to this client (if any)
		chatManager.mutex.Lock()
		chat, exists := chatManager.chats[chatID]
		chatManager.mutex.Unlock()
		if exists {
			chat.MessageMutex.Lock()
			for _, msg := range chat.Messages {
				// Send old messages to client
				s.Emit("message", msg)
			}
			chat.MessageMutex.Unlock()
		}
		return nil
	})

	// On disconnect handler
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		chatID := s.Context().(map[string]string)["chat_id"]
		username := s.Context().(map[string]string)["username"]
		log.Printf("Disconnected: user %s left chat %s: %s", username, chatID, reason)
		s.Leave(chatID)
	})

	// On message from client
	type IncomingMessage struct {
		Content string `json:"content"`
	}

	server.OnEvent("/", "message", func(s socketio.Conn, msg IncomingMessage) {
		ctx := s.Context().(map[string]string)
		chatID := ctx["chat_id"]
		userID := ctx["user_id"]
		username := ctx["username"]
		adminStr := ctx["admin"]
		isAdmin := adminStr == "true" || adminStr == "1"

		if msg.Content == "" {
			return
		}

		// Determine if chat is authorized or guest
		chatManager.mutex.Lock()
		chat, exists := chatManager.chats[chatID]
		chatManager.mutex.Unlock()

		if !exists {
			// If chat does not exist, create it
			authorized := userID != ""
			chat = chatManager.GetOrCreateChat(chatID, authorized)
		}

		message := Message{
			UserID:    userID,
			Username:  username,
			Content:   msg.Content,
			Timestamp: time.Now(),
		}

		// Store the message
		chat.AddMessage(message)

		// Broadcast message to all in chat room
		server.BroadcastToRoom("/", chatID, "message", message)

		log.Printf("[%s] %s: %s", chatID, username, msg.Content)
	})

	// Mount socketio server
	app.Get("/socket.io/*", fiber.WrapHandler(server))

	// HTTP endpoint to create or get chatID for a user (simplified)
	app.Post("/startchat", func(c *fiber.Ctx) error {
		userID, username, authorized := authenticateUser(c)

		var chatID string
		if authorized {
			// for authorized users, chat id can be their user id (or generated from user id)
			chatID = "user_" + userID
		} else {
			// for guests generate temporary chat id, here we keep simple unique ID as timestamp-nano
			chatID = fmt.Sprintf("guest_%d", time.Now().UnixNano())
		}

		// Create chat entry (no messages yet)
		chatManager.GetOrCreateChat(chatID, authorized)

		// Return chat info
		return c.JSON(fiber.Map{
			"chat_id":    chatID,
			"user_id":    userID,
			"username":   username,
			"authorized": authorized,
		})
	})

	// For demo: Admin can see all active chats (in-memory)
	app.Get("/admin/chats", func(c *fiber.Ctx) error {
		// In real app, authenticate admin here
		chatManager.mutex.Lock()
		defer chatManager.mutex.Unlock()

		type ChatInfo struct {
			ID         string    `json:"id"`
			Messages   int       `json:"messages"`
			LastActive time.Time `json:"last_active"`
			Authorized bool      `json:"authorized"`
		}
		chatsInfo := []ChatInfo{}
		for _, chat := range chatManager.chats {
			chatsInfo = append(chatsInfo, ChatInfo{
				ID:         chat.ID,
				Messages:   len(chat.Messages),
				LastActive: chat.LastActive,
				Authorized: chat.Authorized,
			})
		}

		return c.JSON(chatsInfo)
	})

	log.Println("Starting server on :8080")
	log.Fatal(app.Listen(":8080"))
}
