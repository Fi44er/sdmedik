package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

var imageExtensions = map[string]bool{
	"jpg": true, "jpeg": true, "png": true, "gif": true, "webp": true,
}

func IsImageExtension(ext string) bool {
	return imageExtensions[ext]
}

func CompressImageFromMultipart(fileHeader *multipart.FileHeader, outputPath string, quality int) error {
	inputFile, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer inputFile.Close()

	// Читаем содержимое файла в буфер
	buffer, err := io.ReadAll(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Определяем MIME-тип файла
	contentType := http.DetectContentType(buffer)

	// Декодируем изображение
	var img image.Image
	var decodeErr error

	switch contentType {
	case "image/png", "image/jpeg", "image/gif":
		img, _, decodeErr = image.Decode(bytes.NewReader(buffer))
	case "image/webp":
		img, decodeErr = webp.Decode(bytes.NewReader(buffer))
	default:
		return fmt.Errorf("unsupported image format: %s", contentType)
	}

	if decodeErr != nil {
		return fmt.Errorf("failed to decode image: %w", decodeErr)
	}

	// Изменяем размер (опционально)
	img = resize.Resize(1920, 1080, img, resize.Lanczos3)

	// Сохраняем изображение в JPEG
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	options := &jpeg.Options{Quality: quality}
	if err := jpeg.Encode(outputFile, img, options); err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}

func SaveFile(fileHeader *multipart.FileHeader, outputPath string) error {
	inputFile, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}
