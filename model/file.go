package model

import "github.com/tim5wang/selfman/dao/entity"

type File struct {
	entity.File
}

func (u *File) ToEntity() *entity.File {
	return &u.File
}

func (u *File) FromEntity(file *entity.File) {
	u.File = *file
}
