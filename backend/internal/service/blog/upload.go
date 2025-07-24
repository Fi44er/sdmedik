package blog

import (
	"context"
	"path/filepath"

	"github.com/Fi44er/sdmedik/backend/internal/service/blog/utils"
	"github.com/google/uuid"
)

var fileLink = "http://localhost:8080/api/v1/image/"

func (s *service) Upload(ctx context.Context, name string, data []byte) (string, error) {
	expansion := filepath.Ext(name)
	newName := uuid.New().String()

	if err := utils.Upload(&newName, data, expansion); err != nil {
		return "", err
	}

	link := fileLink + newName

	return link, nil
}
