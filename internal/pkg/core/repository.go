package core

import (
	"fmt"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"service/internal/pkg/config"
	"time"
)

type TransactionRepository interface {
	SetTransaction(tx *gorm.DB)
}

func GetIncrementMonthly(model interface{}) int64 {
	var totalData int64
	config.PgSQL.Unscoped().
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
				xtremepkg.LogError(fmt.Sprintf("Truncate invalid: %v", err))
			}
		}
	}
}
