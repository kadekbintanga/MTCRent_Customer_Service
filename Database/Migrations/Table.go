package Migrations

import (
	"Service/App/Model/Activity"
	"Service/App/Model/Testing"
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
		{
			Connection:  Config.PgSQL,
			CreateTable: Testing.TestingSub{},
		},
	}
}
