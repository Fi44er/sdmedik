package filesystem

import "github.com/Fi44er/sdmedik/backend/pkg/logger"

type LocalFileStorage struct {
	basePath string
	logger   *logger.Logger
}

func NewLocalFileStorage(
	basePath string,
) *LocalFileStorage {
	return &LocalFileStorage{
		basePath: basePath,
	}
}

func (s *LocalFileStorage) UploadFile(name string, data []byte) error {
	return nil
}
