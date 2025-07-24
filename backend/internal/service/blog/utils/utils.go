package utils

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

var (
	fileDir = "./image/"
)

func Upload(name *string, data []byte, expansion string) error {
	outputPath := fileDir + *name
	reader := bytes.NewReader(data)

	if img, err := webp.Decode(reader); err == nil {
		return saveAsWebP(img, outputPath)
	}

	reader.Seek(0, io.SeekStart)
	img, _, err := image.Decode(reader)
	if err == nil {
		newPath := replaceExtToWebP(outputPath)
		*name = *name + ".webp"

		return saveAsWebP(img, newPath)
	}

	if err := os.MkdirAll(fileDir, 0755); err != nil {
		fmt.Errorf("failed to create directory: %s", fileDir)
		return err
	}

	*name = *name + "." + expansion
	outputPath = outputPath + "." + expansion
	if os.WriteFile(outputPath, data, 0644) != nil {
		fmt.Errorf("failed to write file: %s", outputPath)
		return err
	}

	return nil
}

func Delete(name string) error {
	filePath := fileDir + name
	if err := os.Remove(filePath); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to delete file %s: %w", name, err)
		}
	}
	return nil
}

func replaceExtToWebP(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimSuffix(path, ext) + ".webp"
}

func saveAsWebP(img image.Image, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	options := webp.Options{Quality: 80, Lossless: false}
	if err := webp.Encode(file, img, &options); err != nil {
		return err
	}

	return nil
}
