package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PromotionType string

const (
	PromotionTypeProductDiscount  PromotionType = "product_discount"  // Скидка на товар
	PromotionTypeCategoryDiscount PromotionType = "category_discount" // Скидка на категорию
	PromotionTypeBuyNGetM         PromotionType = "buy_n_get_m"       // Купи N, получи M
)

type Promotion struct {
	ID          string        `json:"id" gorm:"primaryKey;type:string;"`
	Name        string        `json:"name" gorm:"type:varchar(255);not null"`
	Description string        `json:"description" gorm:"type:text"`
	Type        PromotionType `json:"type" gorm:"type:varchar(50);not null"`
	TargetID    string        `json:"target_id" gorm:"type:string;not null"` // ID категории, группы или товара
	StartDate   time.Time     `json:"start_date" gorm:"type:timestamp;not null"`
	EndDate     time.Time     `json:"end_date" gorm:"type:timestamp;not null"`
	Conditions  []Condition   `json:"conditions" gorm:"foreignKey:PromotionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rewards     []Reward      `json:"rewards" gorm:"foreignKey:PromotionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Promotion) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}

// -------------------------------
type ConditionType string

const (
	ConditionTypeMinQuantity ConditionType = "min_quantity" // Минимальное количество товаров
	ConditionTypeBuyN        ConditionType = "buy_n"        // Купи N товаров
	ConditionTypeGetM        ConditionType = "get_m"        // Получи M товаров
)

type Condition struct {
	ID          string        `json:"id" gorm:"primaryKey;type:string;"`
	PromotionID string        `json:"promotion_id" gorm:"type:string;not null"`
	Type        ConditionType `json:"type" gorm:"type:varchar(50);not null"`
	Value       string        `json:"value" gorm:"type:text"` // Может быть int, float64, string и т.д.
}

func (c *Condition) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

// --------------------------------

type RewardType string

const (
	RewardTypePercentage RewardType = "percentage" // Скидка в процентах
	RewardTypeFixed      RewardType = "fixed"      // Фиксированная скидка
	RewardTypeProduct    RewardType = "product"    // Бесплатный товар
)

type Reward struct {
	ID          string     `json:"id" gorm:"primaryKey;type:string;"`
	PromotionID string     `json:"promotion_id" gorm:"type:string;not null"`
	Type        RewardType `json:"type" gorm:"type:varchar(50);not null"`
	Value       float64    `json:"value" gorm:"type:decimal(10,2);not null"` // Процент, фиксированная сумма или количество товаров
}

func (r *Reward) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()
	return nil
}
