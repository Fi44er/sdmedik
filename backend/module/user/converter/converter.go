package converter

import (
	"github.com/Fi44er/sdmedik/backend/module/user/domain"
	"github.com/Fi44er/sdmedik/backend/module/user/dto"
	"github.com/Fi44er/sdmedik/backend/module/user/model"
)

func ToDomainFromModel(user *model.User) *domain.User {
	return &domain.User{
		ID:          user.ID,
		Email:       user.Email,
		Password:    user.Password,
		FIO:         user.FIO,
		PhoneNumber: user.PhoneNumber,
		Role:        domain.Role(user.Role),
	}
}

func ToModelFromDomain(user *domain.User) *model.User {
	return &model.User{
		ID:          user.ID,
		Email:       user.Email,
		Password:    user.Password,
		FIO:         user.FIO,
		PhoneNumber: user.PhoneNumber,
		Role:        model.Role(user.Role),
	}
}

func ToDomainFromDto(user *dto.UserDTO) *domain.User {
	return &domain.User{
		FIO:         user.FIO,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	}
}

func ToDomainSliceFromModel(models []model.User) []domain.User {
	if models == nil {
		return nil
	}

	domains := make([]domain.User, len(models))
	for i, m := range models {
		domains[i] = *ToDomainFromModel(&m)
	}
	return domains
}

func ToResponseFromDomain(domain *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:          domain.ID,
		Email:       domain.Email,
		FIO:         domain.FIO,
		PhoneNumber: domain.PhoneNumber,
		Role:        string(domain.Role),
	}
}

func ToResponseSliceFromDomain(domains []domain.User) []dto.UserResponse {
	if domains == nil {
		return nil
	}

	response := make([]dto.UserResponse, len(domains))
	for i, d := range domains {
		response[i] = *ToResponseFromDomain(&d)
	}
	return response
}
