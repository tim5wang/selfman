package genid

import (
	"errors"
	"fmt"

	"github.com/tim5wang/selfman/dao/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type IDDao struct {
	Tab *entity.ID
	db  *gorm.DB
}

func NewIDDao(db *gorm.DB) *IDDao {
	return &IDDao{
		Tab: &entity.ID{},
		db:  db,
	}
}

func (d *IDDao) GenID(tab schema.Tabler) (string, error) {
	key := tab.TableName()
	id := entity.ID{}
	var create bool
	err := d.db.Where("key = ?", key).First(&id).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		create = true
	}
	id.Key = key
	id.KeyID++
	if create {
		err = d.db.Create(&id).Error
		return fmt.Sprintf("%d", id.KeyID), err
	}
	err = d.db.Updates(&id).Error
	return fmt.Sprintf("%d", id.KeyID), err
}
