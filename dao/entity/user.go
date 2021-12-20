package entity

type User struct {
	ID             uint64 `json:"id" gorm:"column:id;primary_key;type:bigint(20);comment:自增ID"`
	UserName       string `json:"username" gorm:"column:username;type:char(32);comment:登陆用户名"`
	Password       string `json:"password" gorm:"column:password;type:char(64);comment:密码"`
	NickName       string `json:"nickname" gorm:"column:nickname;type:char(64);comment:昵称"`
	Avator         string `json:"avator" gorm:"column:avator;type:varchar(512);comment:头像"`
	UserType       int    `json:"user_type" gorm:"column:user_type;type:tinyint(2);comment:用户类型"`
	Gender         int    `json:"gender" gorm:"column:gender;type:tinyint(2);comment:性别"`
	Description    string `json:"description" gorm:"column:description;type:varchar(512);comment:签名"`
	Phone          string `json:"phone" gorm:"column:phone;type:char(32);comment:电话"`
	Email          string `json:"email" gorm:"column:email;type:char(32);comment:邮箱"`
	Identify       string `json:"identity" gorm:"column:identify;type:char(32);comment:身份证"`
	CreateTime     int64  `json:"create_time" gorm:"column:create_time;type:bigint(20);comment:创建时间"`
	IdentifyTime   int64  `json:"identify_time" gorm:"column:identify_time;type:bigint(20);comment:认证时间"`
	IdentifyPeriod int64  `json:"identify_period" gorm:"column:identify_period;type:bigint(20);comment:认证有效期限"`
	UserStatus     int    `json:"user_status" gorm:"column:user_status;type:int(10);comment:用户状态"`
	Flag           int64  `json:"flag" gorm:"column:flag;type:bigint(20);comment:保留开关字段"`
	UpdateTime     int64  `json:"update_time" gorm:"column:update_time;type:bigint(20);comment:更新时间"`
}

func (t *User) TableName() string {
	return "tb_user"
}
