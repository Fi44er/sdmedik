package session

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

const sessionContextKey = "session_id"

type ISessionStore interface {
	read(id string) (*Session, error)
	write(session *Session) error
	destroy(id string) error
	gc(idleExpiration, absoluteExpiration time.Duration) error
}

type SessionManager struct {
	store              ISessionStore
	idleExpiration     time.Duration
	absoluteExpiration time.Duration
	cookieName         string
}

func NewSessionManager(
	store ISessionStore,
	gcInterval,
	idleExpiration,
	absoluteExpiration time.Duration,
	cookieName string,
) *SessionManager {

	m := &SessionManager{
		store:              store,
		idleExpiration:     idleExpiration,
		absoluteExpiration: absoluteExpiration,
		cookieName:         cookieName,
	}

	go m.gc(gcInterval)

	return m
}

func (m *SessionManager) gc(d time.Duration) {
	ticker := time.NewTicker(d)

	for range ticker.C {
		m.store.gc(m.idleExpiration, m.absoluteExpiration)
	}
}

func (m *SessionManager) validate(session *Session) bool {
	if time.Since(session.createdAt) > m.absoluteExpiration ||
		time.Since(session.lastActivityAt) > m.idleExpiration {

		// Delete the session from the store
		err := m.store.destroy(session.id)
		if err != nil {
			panic(err)
		}

		return false
	}

	return true
}

func (m *SessionManager) start(ctx *fiber.Ctx) *Session {
	var session *Session
	var err error

	cookie := ctx.Cookies(m.cookieName)
	session, err = m.store.read(cookie)
	if err != nil {
		log.Printf("Failed to read session from store: %v", err)
	}

	if session == nil || !m.validate(session) {
		session = newSession()
	}

	context := context.WithValue(ctx.Context(), sessionContextKey, session)
	ctx.SetUserContext(context)

	return session
}

func (m *SessionManager) save(session *Session) error {
	session.lastActivityAt = time.Now()

	err := m.store.write(session)
	if err != nil {
		return err
	}

	return nil
}

func (m *SessionManager) migrate(session *Session) error {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	err := m.store.destroy(session.id)
	if err != nil {
		return err
	}

	session.id = generateSessionId()

	return nil
}

func (m *SessionManager) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Start the session
		session := m.start(c)

		// Set response headers
		c.Vary("Cookie")
		c.Set("Cache-Control", `no-cache="Set-Cookie"`)

		// Proceed to next handler
		err := c.Next()

		// Save the session
		if saveErr := m.save(session); saveErr != nil {
			log.Printf("Failed to save session: %v", saveErr)
		}

		// Set the session cookie
		c.Cookie(&fiber.Cookie{
			Name:  m.cookieName,
			Value: session.id,
			// Domain:   "test",
			HTTPOnly: true,
			Path:     "/",
			Secure:   false,
			SameSite: "Lax",
			Expires:  time.Now().Add(m.idleExpiration),
			MaxAge:   int(m.idleExpiration / time.Second),
		})

		return err
	}
}

func GetSession(c *fiber.Ctx) *Session {
	session, ok := c.UserContext().Value(sessionContextKey).(*Session)
	if !ok {
		panic("session not found in request context")
	}
	return session
}
