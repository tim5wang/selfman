package model

import "github.com/tim5wang/selfman/dao/entity"

type Doc struct {
	entity.Doc
}

func (u *Doc) ToEntity() *entity.Doc {
	return &u.Doc
}

func (u *Doc) FromEntity(doc *entity.Doc) {
	u.Doc = *doc
}
