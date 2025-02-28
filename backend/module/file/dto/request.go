package dto

import "mime/multipart"

type UploadDTO struct {
	Files []*multipart.FileHeader `json:"files"`
}
