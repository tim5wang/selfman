package entity

type Doc struct {
	BasicField
	DocID    string `json:"doc_id" gorm:"column:doc_id;type:char(32);comment:doc id"`
	Title    string `json:"title" gorm:"column:title;type:varchar(512);comment:标题"`
	Content  string `json:"content" gorm:"column:content;type:text;comment:文章内容"`
	Cas      int64  `json:"cas" gorm:"column:cas;type:bigint(20);comment:compare_and_swap"`
	Comments []byte `json:"comments" gorm:"column:comments;type:blob;comment:评论"`
}

func (t *Doc) TableName() string {
	return "tb_doc"
}
