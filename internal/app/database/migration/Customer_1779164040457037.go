package migration

import (
	"os"

	xtremedb "github.com/globalxtreme/go-core/v2/database"

	"service/internal/pkg/config"
	"service/internal/pkg/model"
)

type Customer_1779164040457037 struct{}

func (Customer_1779164040457037) Reference() string {
	return "Customer_1779164040457037"
}

func (Customer_1779164040457037) Tables() []xtremedb.Table {
	owner := os.Getenv("DB_OWNER")
	return []xtremedb.Table{
		{Connection: config.PgSQL, CreateTable: model.Customer{}, Owner: owner},
	}
}

func (Customer_1779164040457037) Columns() []xtremedb.Column {
	return []xtremedb.Column{}
}
