package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) SaveMessage(ctx context.Context, message *model.Message) error {
	if err := s.repository.SaveMessage(ctx, message); err != nil {
		return err
	}

	templateData := struct {
		Message   string
		ChatID    string
		SenderID  string
		CreatedAt string
		ChatLink  string
	}{
		Message:   message.Message,
		ChatID:    message.ChatID,
		SenderID:  message.SenderID,
		CreatedAt: message.CreatedAt.Format("2006-01-02 15:04:05"),
		ChatLink:  s.config.FrontendURL + "/chat/" + message.ChatID,
	}

	s.mailer.SendMailAsync(
		s.config.MailFrom,
		"New message",
		templateData,
		[]string{"sales@sdmedik.ru", "amanager@sdmedik.ru"},
	)

	return nil
}
