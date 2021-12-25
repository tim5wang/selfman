package entity

type Doc struct {
	BasicField
	Title   string `json:"title" gorm:"column:title;type:varchar(512);comment:标题"`
	Content string `json:"content" gorm:"column:content;type:text;comment:文章内容"`
}

func (t *Doc) TableName() string {
	return "tb_doc"
}
