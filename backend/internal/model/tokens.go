package model

import "time"

type Token struct {
	ID       string    `json:"id"`
	Token    string    `json:"token"`
	UserID   string    `json:"user_id"`
	CreateAt time.Time `json:"create_at"`
}
