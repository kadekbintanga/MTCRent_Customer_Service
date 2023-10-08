package Command

import (
	"fmt"
	"os"
)

type TestCommand struct{}

func (class TestCommand) Handle() {
	storageDir := os.Getenv("STORAGE_DIR") + "/logs/"
	fmt.Println(storageDir)
}
