package index

import (
	"fmt"

	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) AddOrUpdate(data interface{}, docType string) error {
	name, err := utils.FindFieldInObject(data, "Name")
	if err != nil {
		return err
	}

	// Добавляем извлечение артикула
	article, err := utils.FindFieldInObject(data, "Article")
	if err != nil {
		// Если поле не найдено, можно пропустить или установить пустое значение
		article = ""
	}

	id, err := utils.FindFieldInObject(data, "ID")
	if err != nil {
		return err
	}
	strId, err := utils.StringifyID(id)
	if err != nil {
		return err
	}

	doc := map[string]interface{}{
		"Name":    name,
		"Article": article, // Добавляем артикул
		"Type":    docType,
	}

	if err := s.index.Index(strId, doc); err != nil {
		return fmt.Errorf("ошибка при индексации товара с ID %s: %v", strId, err)
	}

	return nil
}
