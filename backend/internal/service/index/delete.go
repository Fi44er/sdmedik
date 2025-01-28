package index

import (
	"fmt"

	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Delete(data interface{}) error {
	id, err := utils.FindFieldInObject(data, "ID")
	if err != nil {
		return err
	}
	strId, err := utils.StringifyID(id)
	if err != nil {
		return err
	}

	if err := s.index.Delete(strId); err != nil {
		return fmt.Errorf("ошибка при удалении документа с ID %s: %v", id, err)
	}

	return nil
}
