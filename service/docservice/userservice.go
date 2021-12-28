package docservice

import (
	"errors"
	"time"

	"github.com/tim5wang/selfman/dao/docdao"
	"github.com/tim5wang/selfman/model"
	"gorm.io/gorm"
)

type DocService struct {
	docDao *docdao.DocDao
}

func NewDocService(docDao *docdao.DocDao) *DocService {
	return &DocService{
		docDao: docDao,
	}
}

func (s *DocService) CreateDoc(doc *model.Doc) (err error) {
	u := doc.ToEntity()
	err = s.docDao.Create(u)
	return
}

func (s *DocService) GetDoc(docID string) (err error, doc *model.Doc) {
	err, u := s.docDao.GetByDocID(docID)
	if err != nil {
		return
	}
	doc = &model.Doc{}
	doc.FromEntity(u)
	return
}

func (s *DocService) SaveDoc(doc *model.Doc) (err error) {
	u := doc.ToEntity()
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	if doc.DocID == "" {
		err = s.docDao.Create(u)
		return
	}
	err, od := s.docDao.GetByDocID(doc.DocID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if od != nil {
		od.UpdateUser = doc.UpdateUser
		od.UpdateTime = time.Now().Unix()
		od.Content = doc.Content
		err = s.docDao.Update(od)
		if err != nil {
			return
		}
	} else {
		err = s.docDao.Create(u)
		if err != nil {
			return
		}
	}
	return
}
