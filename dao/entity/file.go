package entity

type FileEntity struct {
	ID      uint64 `json:"id" gorm:"column:id;primary_key;type:bigint(20);comment:自增ID"`
	Name    string `json:"name" gorm:"column:name;type:varchar(512);comment:文件名称"`
	Size    int64  `json:"size" gorm:"column:size;type:bigint(20);comment:文件大小byte"`
	Mode    uint64 `json:"mode" gorm:"column:mode;type:bigint(20);comment:文件mode信息"`
	ModTime uint64 `json:"mod_time" gorm:"column:mod_time;type:bigint(20);comment:文件修改时间"`
	IsDir   bool   `json:"is_dir" gorm:"column:is_dir;type:tinyint(1);comment:是否时路径"`
	Content []byte `json:"content" gorm:"column:content;type:blob;comment:内容"`
}
