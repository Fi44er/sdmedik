package image

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) DeleteByNames(ctx context.Context, names []string) error {
	var files []string
	for _, name := range names {
		files = append(files, s.config.ImageDir+name)
	}

	if err := utils.DeleteManyFiles(files); err != nil {
		s.logger.Errorf("Error deleting files: %v", err)
		return err
	}

	return nil
}
