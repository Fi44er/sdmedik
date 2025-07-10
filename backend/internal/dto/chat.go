package dto

type AddFragment struct {
	ChatID     string `json:"chat_id" validate:"required"`
	StartMsgID string `json:"start_msg_id" validate:"required"`
}
