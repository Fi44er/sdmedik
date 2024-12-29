package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2/log"
)

func CompressImageFromMultipart(fileHeader *multipart.FileHeader, outputPath string, quality int) error {
	// Открыть файл из *multipart.FileHeader
	inputFile, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer inputFile.Close()

	// Определить формат изображения
	buffer, err := io.ReadAll(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	img, format, err := image.Decode(bytes.NewReader(buffer))
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Создать выходной файл
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Сжать изображение в зависимости от формата
	switch format {
	case "jpeg", "jpg":
		options := &jpeg.Options{Quality: quality}
		err = jpeg.Encode(outputFile, img, options)
	case "png":
		encoder := &png.Encoder{CompressionLevel: png.BestCompression}
		err = encoder.Encode(outputFile, img)
	case "gif":
		err = gif.Encode(outputFile, img, nil)
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}

func DeleteManyFiles(uploadedFiles []string) error {
	for _, uploadedFile := range uploadedFiles {
		if removeErr := os.Remove(uploadedFile); removeErr != nil {
			if os.IsNotExist(removeErr) {
				log.Warnf("File not found, skipping: %s", uploadedFile)
				continue
			}
			log.Errorf("Error deleting file: %v", removeErr)
			return removeErr
		}
	}
	return nil
}
