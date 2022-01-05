package docservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/tim5wang/selfman/dao/docdao"
	"github.com/tim5wang/selfman/dao/entity"
	"github.com/tim5wang/selfman/dao/genid"
	"github.com/tim5wang/selfman/model"
	"gorm.io/gorm"
)

type DocService struct {
	docDao *docdao.DocDao
	idGen  *genid.IDDao
}

func NewDocService(docDao *docdao.DocDao, idGen *genid.IDDao) *DocService {
	return &DocService{
		docDao: docDao,
		idGen:  idGen,
	}
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

func (s *DocService) SaveDoc(doc *model.Doc) (error, *model.Doc) {
	u := doc.ToEntity()
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	if doc.DocID == "" {
		id, err := s.idGen.GenID(u)
		if err != nil {
			err = fmt.Errorf("gen id error %w", err)
			return err, nil
		}
		u.DocID = id
		err, res := s.docDao.Create(u)
		if err != nil {
			return err, nil
		}
		doc.FromEntity(res)
		return err, doc
	}
	err, od := s.docDao.GetByDocID(doc.DocID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err, nil
	}
	if od != nil {
		od.UpdateUser = doc.UpdateUser
		od.UpdateTime = time.Now().Unix()
		od.Content = doc.Content
		err = s.docDao.Update(od)
		if err != nil {
			return err, nil
		}
		doc.FromEntity(od)
	} else {
		u.DocID, err = s.idGen.GenID(u)
		if err != nil {
			err = fmt.Errorf("gen id error %w", err)
			return err, nil
		}
		e, res := s.docDao.Create(u)
		if e != nil {
			return e, doc
		}
		doc.FromEntity(res)
	}
	return nil, doc
}

func (s *DocService) DocList(req *model.DicListReq) (err error, rsp *model.DicListRsp) {
	rsp = &model.DicListRsp{}
	var list []*entity.Doc
	err, rsp.Total, list = s.docDao.GetDocList(req.PageNum, req.PageSize, req.KeyWord)
	if err != nil {
		return
	}
	rsp.List = make([]*model.Doc, 0)
	for _, d := range list {
		doc := &model.Doc{}
		doc.FromEntity(d)
		rsp.List = append(rsp.List, doc)
	}
	return
}
