package user

import (
	"github.com/Fi44er/sdmedik/backend/internal/module/user/dto"
	"github.com/Fi44er/sdmedik/backend/internal/module/user/entity"
)

type Converter struct{}

func (c *Converter) ToEntity(dto *dto.UserDTO) *entity.User {
	return &entity.User{
		Name:         dto.Name,
		Surname:      dto.Surname,
		Patronymic:   dto.Patronymic,
		PhoneNumber:  dto.PhoneNumber,
		Email:        dto.Email,
		PasswordHash: dto.Password,
	}
}

func (c *Converter) ToResponse(entity *entity.User) *dto.UserResponse {
	roles := make([]string, len(entity.Roles))
	for i, r := range entity.Roles {
		roles[i] = r.Name
	}
	return &dto.UserResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Surname:     entity.Surname,
		Patronymic:  entity.Patronymic,
		PhoneNumber: entity.PhoneNumber,
		Email:       entity.Email,
		Role:        roles,
	}
}

func (c *Converter) ToResponseSlice(entities []entity.User) []dto.UserResponse {
	res := make([]dto.UserResponse, len(entities))
	for i, e := range entities {
		res[i] = *c.ToResponse(&e)
	}
	return res
}
