package migration

import (
	xtremedb "github.com/globalxtreme/go-core/v2/database"
	"service/internal/pkg/config"
	"service/internal/pkg/model"
)

func Tables() []xtremedb.Table {
	return []xtremedb.Table{
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
