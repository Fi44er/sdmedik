package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) Update(ctx context.Context, categoryID int, data *dto.UpdateCategory) error {

	if err := s.validator.Struct(data); err != nil {
		return errors.New(400, err.Error())
	}

	// Проверяем, существует ли категория
	existingCategory, err := s.repo.GetByID(ctx, categoryID)
	if err != nil {
		return err
	}
	if existingCategory == nil {
		return errors.New(404, "Category not found")
	}

	tx, err := s.transactionManagerRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			s.logger.Error("Transaction rollback")
			s.transactionManagerRepo.Rollback(tx)
			panic(r) // Переподнимаем панику
		}
	}()

	// Обновляем основную информацию о категории
	existingCategory.Name = data.Name

	if err := s.repo.Update(ctx, existingCategory, tx); err != nil {
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	// Обновляем характеристики
	err = s.characteristicService.Update(ctx, categoryID, data.Characteristics, tx)
	if err != nil {
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	if err := s.transactionManagerRepo.Commit(tx); err != nil {
		return err
	}

	return nil
}
