package utils

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(folderName, fileName string, file multipart.File) bool {
	currentDir, err := os.Getwd()
	path := filepath.Join(currentDir, "uploads", "images", fileName)
	if !CreateDir(path) {
		return false
	}

	dst, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func CreateDir(path string) bool {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return false
		}
	}
	return true
}
