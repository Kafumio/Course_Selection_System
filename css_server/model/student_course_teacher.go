package model

type StudentCourseTeacher struct {
	Sctid int     `json:"sctid"`
	Sid   int     `json:"sid"`
	Cid   int     `json:"cid"`
	Tid   int     `json:"tid"`
	Grade float64 `json:"grade"`
	Term  string  `json:"term"`
}

// TableName  自定义表名
func (this *StudentCourseTeacher) TableName() string {
	return "sct"
}
