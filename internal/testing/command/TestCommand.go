package command

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/spf13/cobra"
	"os"
)

type TestCommand struct{}

func (class *TestCommand) Command(cobraCmd *cobra.Command) {
	cobraCmd.AddCommand(&cobra.Command{
		Use:  "dev-test",
		Long: "Development Test Command",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()

			class.Handle()
		},
	})
}

func (class *TestCommand) Handle() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Working Directory:", dir)
}
