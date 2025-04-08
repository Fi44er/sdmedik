package usecase

import "github.com/Fi44er/sdmedik/backend/pkg/logger"

type IFileRepository interface{}

type IFileStorage interface {
	UploadFiles(name string, data []byte) error
}

type FileUsecase struct {
	repository  IFileRepository
	fileStorage IFileStorage
	logger      *logger.Logger
}

func NewFileUsecase(
	repository IFileRepository,
	fileStorage IFileStorage,
	logger *logger.Logger,
) *FileUsecase {
	return &FileUsecase{
		repository:  repository,
		fileStorage: fileStorage,
		logger:      logger,
	}
}

func (u *FileUsecase) UploadFiles() error {
	return nil
}
