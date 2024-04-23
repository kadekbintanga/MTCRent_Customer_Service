package Runner

import (
	"Service/Config"
	"Service/Database/Migrations"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/database/migration"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "migration",
		Long: "Running Migration",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			Config.InitDB()

			migration.Migrate(Migrations.Tables(), Migrations.Columns())
		},
	})
}
