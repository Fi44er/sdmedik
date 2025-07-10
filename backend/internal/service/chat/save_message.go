package chat

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) SaveMessage(ctx context.Context, message *model.Message) error {
	// 1. Сохраняем сообщение
	if err := s.repository.SaveMessage(ctx, message); err != nil {
		return err
	}

	// 2. Получаем последний фрагмент (включая незавершенные)
	lastFragment, err := s.repository.GetLastChatFragment(ctx, message.ChatID)
	if err != nil {
		return fmt.Errorf("failed to get last fragment: %w", err)
	}

	// 3. Создаем новый фрагмент только если:
	// - нет ни одного фрагмента
	// - последний фрагмент завершен (имеет EndMsgID)
	if lastFragment == nil || lastFragment.EndMsgID != nil {
		dto := &dto.AddFragment{
			ChatID:     message.ChatID,
			StartMsgID: message.ID,
		}
		if err := s.AddFragment(ctx, dto); err != nil {
			return fmt.Errorf("failed to create fragment: %w", err)
		}
	}

	// Отправка email
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
