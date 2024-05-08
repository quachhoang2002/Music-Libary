package file

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// save file to disk, and return file path
func SaveFile(file *multipart.FileHeader, dir string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(filepath.Join(dir, file.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return dst.Name(), nil
}

// get file from file path
func getFilePath(filepath string) string {
	return filepath
}
