package repository

import (
	"github.com/Fi44er/sdmedik/backend/internal/module/user/entity"
	role_model "github.com/Fi44er/sdmedik/backend/internal/module/user/infrastructure/repository/role/model"
	user_model "github.com/Fi44er/sdmedik/backend/internal/module/user/infrastructure/repository/user/model"
)

type Converter struct{}

func (c *Converter) ToModel(entity *entity.User) *user_model.User {
	roles := make([]role_model.Role, len(entity.Roles))
	for i, r := range entity.Roles {
		roles[i] = role_model.Role{ID: r.ID, Name: r.Name}
	}
	return &user_model.User{
		ID:           entity.ID,
		Email:        entity.Email,
		PasswordHash: entity.PasswordHash,
		Name:         entity.Name,
		Surname:      entity.Surname,
		Patronymic:   entity.Patronymic,
		PhoneNumber:  entity.PhoneNumber,
		Roles:        roles,
	}
}

func (c *Converter) ToEntity(model *user_model.User) *entity.User {
	roles := make([]entity.Role, len(model.Roles))
	for i, r := range model.Roles {
		roles[i] = entity.Role{ID: r.ID, Name: r.Name}
	}
	return &entity.User{
		ID:           model.ID,
		Email:        model.Email,
		PasswordHash: model.PasswordHash,
		Name:         model.Name,
		Surname:      model.Surname,
		Patronymic:   model.Patronymic,
		PhoneNumber:  model.PhoneNumber,
		Roles:        roles,
	}
}
