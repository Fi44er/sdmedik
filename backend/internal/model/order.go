package model

import "time"

type Order struct {
	ID             string    `json:"id"`
	FIO            string    `json:"fio"`
	Date           time.Time `json:"date"`
	Article        string    `json:"article"`
	RegionID       int       `json:"region_id"`
	PaymenMethodID int       `json:"payment_method_id"`
	CreateAt       time.Time `json:"create_at"`
}
