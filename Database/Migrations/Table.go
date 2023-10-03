package Migrations

import (
	"Service/App/Models/Activity"
	"Service/App/Models/Testing"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/database/migration"
)

func Tables() []migration.Table {
	return []migration.Table{
		{
			Connection:  Config.PgSQL,
			CreateTable: Activity.Activity{},
		},
		{
			Connection:  Config.PgSQL,
			CreateTable: Testing.Testing{},
		},
	}
}
