package dto

import "mime/multipart"

type Images struct {
	Files []*multipart.FileHeader `json:"file"`
}

type CreateImages struct {
	ProductID string `json:"product_id"`
	Images    Images `json:"images"`
}