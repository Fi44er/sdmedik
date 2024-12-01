package utils

import (
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
)

func HashString(s string) (string, error) {
	h, err := blake2b.New256(nil)
	if err != nil {
		return "", err
	}
	h.Write([]byte(s))
	hashBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString, nil
}
