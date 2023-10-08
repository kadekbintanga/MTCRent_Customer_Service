package Helpers

import (
	"os"
	"strings"
)

func FullPathURL(path string) any {
	baseURL := os.Getenv("STORAGE_GATEWAY_BASE")
	basePath := os.Getenv("STORAGE_PATH")

	return baseURL + basePath + "/" + path
}

func GetRealStoragePath(url string) string {
	baseURL := os.Getenv("STORAGE_GATEWAY_BASE")
	basePath := os.Getenv("STORAGE_PATH")

	return strings.Replace(url, baseURL+basePath+"/", "", -1)
}

func CheckAndCreateDirectory(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		}
	}
}

func CheckAndSetDefaultFile(path string) string {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Getenv("STORAGE_DIR") + "/default/not-found.jpg"
		}
	}

	return path
}
