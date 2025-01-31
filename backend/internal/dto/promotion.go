package dto

import "github.com/Fi44er/sdmedik/backend/internal/model"

type Promotion struct {
	Name        string              `json:"name" validate:"required"`
	Description string              `json:"description"`
	Type        model.PromotionType `json:"type" validate:"required"`
	TargetID    string              `json:"target_id" validate:"required"`
	StartDate   string              `json:"start_date" validate:"required"`
	EndDate     string              `json:"end_date" validate:"required"`
}

type CreatePromotion struct {
	Name         string              `json:"name" validate:"required"`
	Description  string              `json:"description"`
	Type         model.PromotionType `json:"type" validate:"required"`
	TargetID     string              `json:"target_id" validate:"required"`
	GetProductID string              `json:"get_product_id"`
	StartDate    string              `json:"start_date" validate:"required"`
	EndDate      string              `json:"end_date" validate:"required"`
	Condition    CreateCondition     `json:"condition"`
	Reward       CreateReward        `json:"reward"`
}

type CreateCondition struct {
	Type  model.ConditionType `json:"type" validate:"required"`
	Value string              `json:"value" validate:"required"`
}

type CreateReward struct {
	Type  model.RewardType `json:"type" validate:"required"`
	Value float64          `json:"value" validate:"required"`
}
