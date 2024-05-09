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
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filePath := filepath.Join(dir, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return "", err
			}
		}

		dst, err = os.Create(filePath)
		if err != nil {
			return "", err
		}
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// get file from file path
func getFilePath(filepath string) string {
	return filepath
}
