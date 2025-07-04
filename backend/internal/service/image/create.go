package image

import (
	"context"
	"fmt"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/google/uuid"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"gorm.io/gorm"
)

func (s *service) CreateMany(ctx context.Context, dto *dto.CreateImages, tx *gorm.DB) error {
	var limit int
	existImages := new([]model.Image)
	var err error
	var uploadedFiles []string
	var images []model.Image

	if dto.ProductID != "" {
		existImages, err = s.repo.GetByID(ctx, &dto.ProductID, nil, tx)
		limit = 8
	} else if dto.CategoryID != 0 {
		existImages, err = s.repo.GetByID(ctx, nil, &dto.CategoryID, tx)
		limit = 1
	} else {
		return errors.New(400, "Either ProductID or CategoryID must be provided")
	}

	if err != nil {
		return err
	}

	if len(*existImages)+len(dto.Images.Files) > limit {
		errMsg := fmt.Sprintf("Limit of %d images exceeded", limit)
		return errors.New(400, errMsg)
	}

	for _, image := range dto.Images.Files {
		lastDot := strings.LastIndex(image.Filename, ".")
		name := uuid.New().String() + image.Filename[lastDot:]
		fullPath := s.config.ImageDir + name
		if err := utils.CompressImageFromMultipart(image, fullPath, 40); err != nil {
			if err := utils.DeleteManyFiles(uploadedFiles); err != nil {
				s.logger.Errorf("Failed to remove file %s: %v", uploadedFiles, err)
			}
			s.logger.Errorf("%v", err)
			return err
		}

		uploadedFiles = append(uploadedFiles, fullPath)

		newImage := model.Image{
			Name: name,
		}
		if dto.ProductID != "" {
			newImage.ProductID = &dto.ProductID
		} else if dto.CategoryID != 0 {
			newImage.CategoryID = &dto.CategoryID
		}
		images = append(images, newImage)
	}

	if err := s.repo.CreateMany(ctx, &images, tx); err != nil {
		if err := utils.DeleteManyFiles(uploadedFiles); err != nil {
			s.logger.Errorf("Failed to remove file %s: %v", uploadedFiles, err)
		}
		return err
	}

	return nil
}
