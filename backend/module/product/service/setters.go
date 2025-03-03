package service

import (
	"context"
	"mime/multipart"

	"github.com/Fi44er/sdmedik/backend/converter"
	"github.com/Fi44er/sdmedik/backend/module/product/domain"
	"github.com/Fi44er/sdmedik/backend/module/product/dto"
)

func (s *ProductService) Create(ctx context.Context, productDomain *domain.Product, files []*multipart.FileHeader) error {
	categories, err := s.categoryRepo.GetByIDs(ctx, productDomain.CategoryIDs)
	if err != nil {
		s.logger.Errorf("Failed to get categories: %v", err)
		return err
	}

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

	if err := s.repo.Create(ctx, productDomain, tx); err != nil {
		s.logger.Errorf("Failed to create product: %v", err)
		s.transactionManagerRepo.Rollback(ctx, tx)
		return err
	}

	productCategryDomains := make([]domain.ProductCategory, len(categories))
	for i, category := range categories {
		productCategryDomains[i] = domain.ProductCategory{
			CategoryID: category.ID,
			ProductID:  productDomain.ID,
		}
	}

	s.logger.Infof("productDomain: %+v", categories)

	if err := s.repo.CreateProductCategoies(ctx, productCategryDomains, tx); err != nil {
		s.logger.Errorf("Failed to create product categories: %v", err)
		s.transactionManagerRepo.Rollback(ctx, tx)
		return err
	}

	createFileDTO := dto.CreateFileDTO{
		OwnerID:   productDomain.ID,
		OwnerType: "product",
	}

	fileDomain := converter.CreateProductFileToFileDomain(&createFileDTO)
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

	s.logger.Info("Product created successfully with files")
	return nil
}
