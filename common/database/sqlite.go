package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateSQLite(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()
	return db
}
