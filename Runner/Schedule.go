package Runner

import (
	"Service/App/Console"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/console"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "schedule",
		Long: "Running Schedule",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			Config.InitDB()

			console.Schedules(Console.RegisterSchedule)
		},
	})
}
