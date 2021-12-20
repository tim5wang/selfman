package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func StartTransaction(db *gorm.DB) (func(error), *gorm.DB) {
	tx := db.Begin()
	rollbackWhenError := func(err error) {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
	return rollbackWhenError, tx
}

func GetTable(db *gorm.DB, tab schema.Tabler, tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 && tx[0] != nil {
		return tx[0].Table(tab.TableName())
	}
	return db.Table(tab.TableName())
}
