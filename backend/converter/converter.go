package converter

import (
	auth_dto "github.com/Fi44er/sdmedik/backend/module/auth/dto"
	category_dto "github.com/Fi44er/sdmedik/backend/module/category/dto"
	file_domain "github.com/Fi44er/sdmedik/backend/module/file/domain"
	product_dto "github.com/Fi44er/sdmedik/backend/module/product/dto"
	user_domain "github.com/Fi44er/sdmedik/backend/module/user/domain"
)

func ConvertRegisterDtoToUserDomain(dto *auth_dto.RegisterDTO) *user_domain.User {
	return &user_domain.User{
		Email:       dto.Email,
		Password:    dto.Password,
		FIO:         dto.FIO,
		PhoneNumber: dto.PhoneNumber,
	}
}

func CreateProductFileToFileDomain(dto *product_dto.CreateFileDTO) *file_domain.File {
	return &file_domain.File{
		OwnerID:   dto.OwnerID,
		OwnerType: dto.OwnerType,
	}
}

func CreateCategoryFileToFileDomain(dto *category_dto.CreateFileDTO) *file_domain.File {
	return &file_domain.File{
		OwnerID:   dto.OwnerID,
		OwnerType: dto.OwnerType,
	}
}
