package provider

import "github.com/Fi44er/sdmedik/backend/internal/api/chat"

type ChatProvider struct {
	chatImpl *chat.Implementation
}

func NewChatProvider() *ChatProvider {
	return &ChatProvider{}
}
func (p *ChatProvider) ChatImpl() *chat.Implementation {
	if p.chatImpl == nil {
		p.chatImpl = chat.NewImplementation()
	}
	return p.chatImpl
}
