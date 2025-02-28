package service

import (
	"context"
	"mime/multipart"

	"github.com/Fi44er/sdmedik/backend/converter"
	"github.com/Fi44er/sdmedik/backend/module/category/domain"
	"github.com/Fi44er/sdmedik/backend/module/category/dto"
)

func (s *CategoryService) Create(ctx context.Context, categoryDomain *domain.Category, files []*multipart.FileHeader) error {
	tx, newCtx, err := s.transactionManagerRepo.Begin(ctx) // Начинаем транзакцию
	if err != nil {
		s.logger.Errorf("Error beginning transaction: %v", err)
		return err
	}
	ctx = newCtx

	defer func() {
		if r := recover(); r != nil {
			s.logger.Error("Transaction rollback due to panic")
			s.transactionManagerRepo.Rollback(ctx, tx)
			panic(r)
		}
	}()

	if err := s.repo.Create(ctx, categoryDomain, tx); err != nil {
		s.logger.Errorf("Failed to create category: %v", err)
		s.transactionManagerRepo.Rollback(ctx, tx)
		return err
	}

	createFileDTO := dto.CreateFileDTO{
		OwnerID:   categoryDomain.ID,
		OwnerType: "category",
	}

	fileDomain := converter.CreateCategoryFileToFileDomain(&createFileDTO)
	_, err = s.fileServ.CreateMany(ctx, fileDomain, files)
	if err != nil {
		s.logger.Errorf("Failed to upload files: %v", err)
		s.transactionManagerRepo.Rollback(ctx, tx)
		return err
	}

	if err := s.transactionManagerRepo.Commit(ctx, tx); err != nil {
		s.logger.Errorf("Failed to commit transaction: %v", err)
		return err
	}

	return nil
}
