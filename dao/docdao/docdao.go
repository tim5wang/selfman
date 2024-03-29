package docdao

import (
	"github.com/tim5wang/selfman/common/database"
	"github.com/tim5wang/selfman/dao/entity"
	"gorm.io/gorm"
)

type DocDao struct {
	Tab *entity.Doc
	db  *gorm.DB
}

func NewDocDao(db *gorm.DB) *DocDao {
	return &DocDao{
		Tab: &entity.Doc{},
		db:  db,
	}
}

func (d *DocDao) Create(doc *entity.Doc, tx ...*gorm.DB) (err error, r *entity.Doc) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	res := tab.Create(doc)
	err = res.Error
	r = doc
	return
}
func (d *DocDao) Update(doc *entity.Doc, tx ...*gorm.DB) (err error) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	res := tab.Where("doc_id = ?", doc.DocID).Updates(doc)
	err = res.Error
	return
}

func (d *DocDao) GetByDocID(docID string, tx ...*gorm.DB) (err error, doc *entity.Doc) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	err = tab.Where("doc_id = ?", docID).First(&doc).Error
	return
}

func (d *DocDao) GetDocList(page, size int, keyword string, tx ...*gorm.DB) (err error, total int64, docs []*entity.Doc) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	if keyword != "" {
		like := "%s" + keyword + "%s"
		tab = tab.Where("title like ? OR content like ?", like, like)
	}
	err = tab.Count(&total).Error
	if page != 0 && size != 0 {
		err = tab.Offset((page - 1) * size).Limit(size).Find(&docs).Error
	} else {
		err = tab.Find(&docs).Error
	}
	return
}
