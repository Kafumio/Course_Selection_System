package model

type CourseTeacher struct {
	Ctid int    `json:"ctid"`
	Cid  int    `json:"cid"`
	Tid  int    `json:"tid"`
	Term string `json:"term"`
}

// TableName  自定义表名
func (this *CourseTeacher) TableName() string {
	return "ct"
}
