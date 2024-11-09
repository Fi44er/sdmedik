package utils

import (
	"encoding/json"
	"os"
)

func SaveToFile[T any](products []T, filename string) error {
	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}
