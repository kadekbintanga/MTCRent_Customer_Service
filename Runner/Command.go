package Runner

import (
	"Service/App/Console"
)

func init() {
	Console.RegisterCommand(rootCmd)
}
