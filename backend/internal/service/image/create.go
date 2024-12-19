package image

import (
	"context"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/google/uuid"

	// "github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"gorm.io/gorm"
)

func (s *service) CreateMany(ctx context.Context, dto *dto.CreateImages, tx *gorm.DB) error {
	existImages, err := s.repo.GetByProductID(ctx, dto.ProductID, tx)
	if err != nil {
		return err
	}

	if len(existImages)+len(dto.Images.Files) == 5 {
		return errors.New(400, "Product already has 5 images")
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

		images = append(images, model.Image{
			ProductID: dto.ProductID,
			Name:      name,
		})
	}

	if err := s.repo.CreateMany(ctx, &images, tx); err != nil {
		return err
	}

	return nil
}
