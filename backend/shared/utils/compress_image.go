package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"os"

	"github.com/nfnt/resize"
)

var imageExtensions = map[string]bool{
	"jpg": true, "jpeg": true, "png": true, "gif": true,
	"bmp": true, "tiff": true, "tif": true, "webp": true,
	"heif": true, "heic": true, "avif": true,
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

	buffer, err := io.ReadAll(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	img, _, err := image.Decode(bytes.NewReader(buffer))
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	img = resize.Resize(1920, 1080, img, resize.Lanczos3)

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
