package model

type ByPage struct {
	PageSize int    `json:"page_size,omitempty"`
	PageNum  int    `json:"page_num,omitempty"`
	Total    int64  `json:"total,omitempty"`
	KeyWord  string `json:"key_word,omitempty"`
}
