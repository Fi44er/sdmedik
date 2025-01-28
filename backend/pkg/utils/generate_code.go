package utils

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)

	for i := range code {
		code[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(code)
}
