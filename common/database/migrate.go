package database

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Migration struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) *Migration {
	return &Migration{db: db}
}

func (m *Migration) Migrate(tabs ...schema.Tabler) (err error) {
	for _, tab := range tabs {
		err = m.db.Migrator().DropTable(tab)
		if err != nil {
			return
		}
		err = m.db.AutoMigrate(tab)
		if err != nil {
			return
		}
		ok := m.db.Migrator().HasTable(tab)
		if !ok {
			err = fmt.Errorf("create %s faild", tab.TableName())
			return
		}
	}
	return
}
