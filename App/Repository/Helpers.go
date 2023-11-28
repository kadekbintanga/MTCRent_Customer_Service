package Repository

import (
	"Service/Config"
	"fmt"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func GetIncrementMonthly(model interface{}) int64 {
	var totalData int64
	Config.PgSQL.Unscoped().
		Where("EXTRACT(MONTH FROM \"createdAt\") = ?", time.Now().Month()).
		Model(&model).
		Count(&totalData)

	return totalData + 1
}

func Truncate(db *gorm.DB, tables ...schema.Tabler) {
	if len(tables) > 0 {
		for _, table := range tables {
			err := db.Exec(fmt.Sprintf("truncate table %s restart identity cascade", table.TableName()))
			if err != nil {
				xtremelog.Error(fmt.Sprintf("Truncate invalid: %v", err))
			}
		}
	}
}
