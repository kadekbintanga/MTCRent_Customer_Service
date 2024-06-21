package migration

import (
	"github.com/globalxtreme/gobaseconf/database/migration"
	"service/internal/pkg/config"
	"service/internal/pkg/model"
)

func Tables() []migration.Table {
	return []migration.Table{
		{
			Connection:  config.PgSQL,
			CreateTable: model.Activity{},
		},
		{
			Connection:  config.PgSQL,
			CreateTable: model.Testing{},
		},
		{
			Connection:  config.PgSQL,
			CreateTable: model.TestingSub{},
		},
	}
}
