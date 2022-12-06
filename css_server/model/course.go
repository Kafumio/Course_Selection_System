package model

type Course struct {
	Cid     int    `json:"cid"`
	Cname   string `json:"cname"`
	Ccredit int    `json:"ccredit"`
}

// TableName  自定义表名
func (this *Course) TableName() string {
	return "c"
}
