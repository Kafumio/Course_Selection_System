package model

type CourseTeacherInfo struct {
	Cid   int     `json:"cid"`
	Tid   int     `json:"tid"`
	Cname string  `json:"cname"`
	Tname string  `json:"tname"`
	Ccredit int   `json:"ccredit"`
	Grade float64 `json:"grade"`
}
