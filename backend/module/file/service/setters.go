package service

import (
	"context"
	"mime/multipart"
	"strings"

	"os"

	"github.com/Fi44er/sdmedik/backend/module/file/domain"
	"github.com/Fi44er/sdmedik/backend/shared/utils"
	"github.com/google/uuid"
)

func (s *FileService) CreateMany(ctx context.Context, fileDomain *domain.File, files []*multipart.FileHeader) ([]string, error) {
	var uploadedFiles []string

	tx, err := s.transactionManagerRepo.GetTransaction(ctx)
	if err != nil {
		s.logger.Errorf("Failed to get transaction: %v", err)
		return nil, err
	}

	for _, file := range files {
		lastDot := strings.LastIndex(file.Filename, ".")
		ext := strings.ToLower(file.Filename[lastDot+1:])
		name := uuid.New().String() + file.Filename[lastDot:]
		fullPath := s.config.ImageDir + name

		if utils.IsImageExtension(ext) {
			if err := utils.CompressImageFromMultipart(file, fullPath, 40); err != nil {
				s.logger.Errorf("Failed to compress image: %v", err)
				s.cleanupFiles(uploadedFiles)
				return nil, err
			}
		} else {
			if err := utils.SaveFile(file, fullPath); err != nil {
				s.logger.Errorf("Failed to save file: %v", err)
				s.cleanupFiles(uploadedFiles)
				return nil, err
			}
		}

		uploadedFiles = append(uploadedFiles, name)
	}

	var fileDomains []domain.File
	for _, file := range uploadedFiles {
		fileDomains = append(fileDomains, domain.File{
			Name:      file,
			OwnerID:   fileDomain.OwnerID,
			OwnerType: fileDomain.OwnerType,
		})
	}

	if err := s.repository.CreateMany(ctx, fileDomains, tx); err != nil {
		return nil, err
	}

	return uploadedFiles, nil
}

func (s *FileService) cleanupFiles(files []string) {
	for _, file := range files {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			s.logger.Errorf("Failed to remove file %s: %v", file, err)
		}
	}
}
