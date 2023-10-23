package Command

import (
	"fmt"
	"os"
)

type TestCommand struct{}

func (class TestCommand) Handle() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Working Directory:", dir)
}
