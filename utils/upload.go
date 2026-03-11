package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// SaveUploadedFile saves a multipart file to the given directory and returns the saved filename
func SaveUploadedFile(file *multipart.FileHeader, destDir string) (string, error) {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	ext := filepath.Ext(file.Filename)
	newName := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102150405"), uuid.New().String(), ext)
	destPath := filepath.Join(destDir, newName)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	return newName, nil
}

// DeleteFile removes a file from the filesystem
func DeleteFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil // File doesn't exist, not an error
	}
	return os.Remove(filePath)
}

// IsAllowedImageType checks whether a file has an allowed image extension
func IsAllowedImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}
	return allowed[ext]
}

// IsAllowedDocumentType checks whether a file has an allowed document extension
func IsAllowedDocumentType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	allowed := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
		".xls":  true,
		".xlsx": true,
	}
	return allowed[ext]
}
