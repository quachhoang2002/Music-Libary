package file

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// save file to disk, and return file path
// if not exist will create new file
func SaveFile(file *multipart.FileHeader, dir string) (string, error) {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dstPath := filepath.Join(dir, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return dst.Name(), nil
}

// get file from file path
func getFilePath(filepath string) string {
	return filepath
}
