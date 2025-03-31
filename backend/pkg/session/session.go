package session

import "time"

type Session struct {
	createdAt      time.Time
	lastActivityAt time.Time
	id             string
	data           map[string]any
}

func newSession() *Session {
	return &Session{
		id:             generateSessionId(),
		data:           make(map[string]any),
		createdAt:      time.Now(),
		lastActivityAt: time.Now(),
	}
}

func (s *Session) Get(key string) any {
	return s.data[key]
}

func (s *Session) Put(key string, value any) {
	s.data[key] = value
}

func (s *Session) Delete(key string) {
	delete(s.data, key)
}
