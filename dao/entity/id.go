package entity

type ID struct {
	BasicField
	Key   string `json:"key" gorm:"column:key;type:char(32);comment:属性"`
	KeyID uint64 `json:"key_id" gorm:"column:id;type:bigint(20)"`
}

func (t *ID) TableName() string {
	return "tb_id"
}
