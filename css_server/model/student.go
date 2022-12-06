package model

type Student struct {
	Sid      int    `json:"sid"`
	Sname    string `json:"sname"`
	Password string `json:"password"`
}

// TableName  自定义表名
func (this *Student) TableName() string {
	return "s"
}
