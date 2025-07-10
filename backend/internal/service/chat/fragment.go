package chat

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

var FragmentColors = []string{
	"#FF0000",
	"#00FF00",
	"#0000FF",
	"#FFFF00",
	"#FF00FF",
	"#00FFFF",
	"#FF8000",
	"#8000FF",
}

func (s *service) AddFragment(ctx context.Context, data *dto.AddFragment) error {
	// Проверяем наличие активного фрагмента
	activeFragment, err := s.repository.GetActiveFragment(ctx, data.ChatID)
	if err != nil {
		return fmt.Errorf("failed to check active fragments: %w", err)
	}

	if activeFragment != nil {
		return nil // Активный фрагмент уже существует
	}

	var fragmentModel model.Fragment
	if err := utils.DtoToModel(data, &fragmentModel); err != nil {
		return err
	}

	// Устанавливаем цвет фрагмента
	lastFragment, err := s.repository.GetLastChatFragment(ctx, data.ChatID)
	if err != nil {
		return fmt.Errorf("failed to get last fragment: %w", err)
	}

	fragmentModel.Color = FragmentColors[0]
	if lastFragment != nil {
		for i := 0; i < len(FragmentColors); i++ {
			if lastFragment.Color == FragmentColors[i] {
				fragmentModel.Color = FragmentColors[(i+1)%len(FragmentColors)]
				break
			}
		}
	}

	return s.repository.CreateFragment(ctx, &fragmentModel)
}

func (s *service) AddEndMsgID(ctx context.Context, chatID string) error {
	fragment, err := s.repository.GetActiveFragment(ctx, chatID)
	if err != nil {
		return err
	}

	msgs, err := s.repository.GetMessagesInFragment(ctx, *fragment)
	if err != nil {
		return err
	}

	if len(msgs) != 0 {
		fragment.EndMsgID = &msgs[len(msgs)-1].ID
	}

	return s.repository.UpdateFragment(ctx, fragment)
}
