package Resizing

import (
	"fmt"
	"os"
	"strings"
)

type Resizing struct {
}

func (resize Resizing) SetFilePath(path string) (string, string, string) {
	baseDir, _ := os.Getwd()
	storagePath := os.Getenv("STORAGE_DIR")

	paths := strings.Split(path, "/")
	filename := paths[len(paths)-1]

	paths = paths[:len(paths)-1]
	newPath := strings.Join(paths, "/")

	return fmt.Sprintf("%s/%s/%s", baseDir, storagePath, path),
		fmt.Sprintf("%s/%s/%s/", baseDir, storagePath, newPath),
		filename
}
