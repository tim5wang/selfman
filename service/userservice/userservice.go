package userservice

import (
	"github.com/tim5wang/selfman/dao/userdao"
	"github.com/tim5wang/selfman/model"
)

type UserService struct {
	userDao *userdao.UserDao
}

func NewUserService(userDao *userdao.UserDao) *UserService {
	return &UserService{
		userDao: userDao,
	}
}

func (s *UserService) CreateUser(user *model.User) (err error) {
	u := user.ToEntity()
	err = s.userDao.Create(u)
	return
}

func (s *UserService) GetUser(username string) (err error, user *model.User) {
	err, u := s.userDao.GetByUsername(username)
	if err != nil {
		return
	}
	user = &model.User{}
	user.FromEntity(u)
	return
}
