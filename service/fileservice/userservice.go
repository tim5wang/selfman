package fileservice

import (
	"github.com/tim5wang/selfman/dao/filedao"
	"github.com/tim5wang/selfman/model"
)

type FileService struct {
	fileDao *filedao.FileDao
}

func NewFileService(fileDao *filedao.FileDao) *FileService {
	return &FileService{
		fileDao: fileDao,
	}
}

func (s *FileService) CreateFile(file *model.File) (err error) {
	u := file.ToEntity()
	err = s.fileDao.Create(u)
	return
}

func (s *FileService) GetFile(filename string) (err error, file *model.File) {
	err, u := s.fileDao.GetByFilename(filename)
	if err != nil {
		return
	}
	file = &model.File{}
	file.FromEntity(u)
	return
}
