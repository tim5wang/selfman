package database

import (
	"fmt"

	"github.com/tim5wang/selfman/common/configservice"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Migration struct {
	db   *gorm.DB
	conf *configservice.ConfigService
}

func NewMigration(db *gorm.DB, c *configservice.ConfigService) *Migration {
	return &Migration{db: db, conf: c}
}

func (m *Migration) Migrate(tabs ...schema.Tabler) (err error) {
	for _, tab := range tabs {
		has := m.db.Migrator().HasTable(tab)
		if has && !m.conf.GetBool("gorm.migrate.drop") {
			continue
		}
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
