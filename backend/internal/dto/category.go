package dto

type CreateCategory struct {
	Name            string                            `json:"name" validate:"required"`
	Characteristics []CharacteristicWithoutCategoryID `json:"characteristics" validate:"dive"`
}

type CategoryWithoutCharacteristics struct {
	Name string `json:"name"`
}
