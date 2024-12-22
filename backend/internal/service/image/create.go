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
	var existImages []model.Image
	var err error

	// Проверяем, передан ли ProductID или CategoryID
	if dto.ProductID != "" {
		existImages, err = s.repo.GetByID(ctx, &dto.ProductID, nil, tx)
		limit = 5
	} else if dto.CategoryID != 0 {
		existImages, err = s.repo.GetByID(ctx, nil, &dto.CategoryID, tx)
		limit = 1
	} else {
		return errors.New(400, "Either ProductID or CategoryID must be provided")
	}

	if err != nil {
		return err
	}

	// Проверяем, не превышает ли количество изображений 5
	if len(existImages)+len(dto.Images.Files) > limit {
		errMsg := fmt.Sprintf("Limit of %d images exceeded", limit)
		return errors.New(400, errMsg)
	}

	path := s.config.ImageDir
	var images []model.Image
	for _, image := range dto.Images.Files {
		lastDot := strings.LastIndex(image.Filename, ".")
		name := uuid.New().String() + image.Filename[lastDot:]
		if err := utils.CompressImageFromMultipart(image, path+name, 40); err != nil {
			s.logger.Errorf("%v", err)
			return err
		}

		// Создаем изображение с соответствующим идентификатором
		newImage := model.Image{
			Name: name,
		}
		if dto.ProductID != "" {
			newImage.ProductID = dto.ProductID
		} else if dto.CategoryID != 0 {
			newImage.CategoryID = dto.CategoryID
		}
		images = append(images, newImage)
	}

	if err := s.repo.CreateMany(ctx, &images, tx); err != nil {
		return err
	}

	return nil
}
