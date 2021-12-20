package model

import "github.com/tim5wang/selfman/dao/entity"

type User struct {
	entity.User
}

func (u *User) ToEntity() *entity.User {
	return &u.User
}

func (u *User) FromEntity(user *entity.User) {
	u.User = *user
}
