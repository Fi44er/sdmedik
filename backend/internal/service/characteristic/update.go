package characteristic

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

func (s *service) Update(ctx context.Context, categoryID int, characteristics []dto.UpdateCharacteristic, tx *gorm.DB) error {
	existingCharacteristics, err := s.repo.GetByCategoryID(ctx, categoryID)
	if err != nil {
		return err
	}

	existingMap := make(map[int]model.Characteristic)
	for _, characteristic := range *existingCharacteristics {
		existingMap[characteristic.ID] = characteristic
	}

	var newCharacteristics []model.Characteristic
	var toDelete []int

	// Определяем, какие характеристики нужно обновить или добавить
	for _, characteristic := range characteristics {
		if characteristic.ID == 0 {
			// Новая характеристика
			newCharacteristics = append(newCharacteristics, model.Characteristic{
				Name:       characteristic.Name,
				CategoryID: categoryID,
				DataType:   model.Type(characteristic.DataType),
			})
		} else {
			// Обновление существующей характеристики
			if existingChar, exists := existingMap[characteristic.ID]; exists {
				existingChar.Name = characteristic.Name
				existingChar.DataType = model.Type(characteristic.DataType)
				if err := s.repo.Update(ctx, &existingChar, tx); err != nil {
					return err
				}
				delete(existingMap, characteristic.ID)
			}
		}
	}

	// Оставшиеся в `existingMap` характеристики нужно удалить
	for id := range existingMap {
		toDelete = append(toDelete, id)
	}

	// Удаляем характеристики
	if len(toDelete) > 0 {
		if err := s.repo.DeleteMany(ctx, toDelete, tx); err != nil {
			return err
		}
	}

	// Добавляем новые характеристики
	if len(newCharacteristics) > 0 {
		if err := s.repo.CreateMany(ctx, &newCharacteristics, tx); err != nil {
			return err
		}
	}

	return nil
}
