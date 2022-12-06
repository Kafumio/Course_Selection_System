package model

type Teacher struct {
	Tid      int    `json:"tid"`
	Tname    string `json:"tname"`
	Password string `json:"password"`
}

// TableName  自定义表名
func (this *Teacher) TableName() string {
	return "t"
}
