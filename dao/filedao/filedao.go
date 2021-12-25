package filedao

import (
	"github.com/tim5wang/selfman/common/database"
	"github.com/tim5wang/selfman/dao/entity"
	"gorm.io/gorm"
)

type FileDao struct {
	Tab *entity.File
	db  *gorm.DB
}

func NewFileDao(db *gorm.DB) *FileDao {
	return &FileDao{
		Tab: &entity.File{},
		db:  db,
	}
}

func (d *FileDao) Create(file *entity.File, tx ...*gorm.DB) (err error) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	res := tab.Create(file)
	err = res.Error
	return
}

func (d *FileDao) GetByFilename(filename string, tx ...*gorm.DB) (err error, file *entity.File) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	err = tab.Where("name = ?", filename).First(&file).Error
	return
}
