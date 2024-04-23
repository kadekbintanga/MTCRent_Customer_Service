package Runner

import (
	"Service/Config"
	"Service/Database/Seeder"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "seeder",
		Long: "Running Seeder",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			Config.InitDB()

			Seeder.Run()
		},
	})
}
