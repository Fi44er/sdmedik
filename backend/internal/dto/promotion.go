package dto

import "github.com/Fi44er/sdmedik/backend/internal/model"

type CreatePromotion struct {
	Name        string              `json:"name" validate:"required"`
	Description string              `json:"description"`
	Type        model.PromotionType `json:"type" validate:"required"`
	TargetID    string              `json:"target_id" validate:"required"`
	StartDate   string              `json:"start_date" validate:"required"`
	EndDate     string              `json:"end_date" validate:"required"`
	Conditions  []CreateCondition   `json:"conditions" validate:"dive"`
	Rewards     []CreateReward      `json:"rewards" validate:"dive"`
}

type CreateCondition struct {
	// PromotionID string `json:"promotion_id" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type CreateReward struct {
	// PromotionID string  `json:"promotion_id" validate:"required"`
	Type  string  `json:"type" validate:"required"`
	Value float64 `json:"value" validate:"required"`
}
