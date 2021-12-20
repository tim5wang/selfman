package userdao

import (
	"github.com/tim5wang/selfman/common/database"
	"github.com/tim5wang/selfman/dao/entity"
	"gorm.io/gorm"
)

type UserDao struct {
	Tab *entity.User
	db  *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		Tab: &entity.User{},
		db:  db,
	}
}

func (d *UserDao) Create(user *entity.User, tx ...*gorm.DB) (err error) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	res := tab.Create(user)
	err = res.Error
	return
}

func (d *UserDao) GetByUsername(username string, tx ...*gorm.DB) (err error, user *entity.User) {
	tab := database.GetTable(d.db, d.Tab, tx...)
	err = tab.Where("username = ?", username).First(&user).Error
	return
}
