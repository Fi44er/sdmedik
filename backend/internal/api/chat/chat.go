package chat

import "sync"

type Implementation struct {
	connections map[string]string // Мапа для хранения подключений (ключ — идентификатор пользователя)
	mu          sync.RWMutex
	userToChat  map[string]string   // ID пользователя, ID чата
	chatToUsers map[string][]string // ID чата, участники
}

func NewImplementation() *Implementation {
	return &Implementation{
		connections: make(map[string]string),
		mu:          sync.RWMutex{},
		userToChat:  make(map[string]string),
		chatToUsers: make(map[string][]string),
	}
}
