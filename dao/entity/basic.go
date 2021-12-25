package entity

type BasicField struct {
	ID         uint64 `json:"id" gorm:"column:id;primary_key;type:bigint(20);comment:自增ID"`
	CreateTime uint64 `json:"create_time" gorm:"column:create_time;type:bigint(20);comment:创建时间"`
	UpdateTime uint64 `json:"update_time" gorm:"column:update_time;type:bigint(20);comment:更新时间"`
	CreateUser string `json:"create_user" gorm:"column:create_user;type:char(32);comment:创建者"`
	UpdateUser string `json:"update_user" gorm:"column:update_user;type:char(32);comment:更新着"`
}
