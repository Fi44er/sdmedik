package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) Delete(ctx context.Context, id string) error {
	names := []string{}

	product, err := s.repo.Get(ctx, dto.ProductSearchCriteria{ID: id})
	if err != nil {
		return err
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

	if err := s.repo.Delete(ctx, id); err != nil {
		if err.Error() == "Product not found" {
			s.transactionManagerRepo.Rollback(tx)
			return errors.New(404, "Product not found")
		}
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	for _, image := range product[0].Images {
		names = append(names, image.Name)
	}

	if err := s.imageService.DeleteByNames(ctx, names); err != nil {
		s.logger.Errorf("Error deleting files: %v", err)
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	if err := s.transactionManagerRepo.Commit(tx); err != nil {
		return err
	}

	return nil
}
