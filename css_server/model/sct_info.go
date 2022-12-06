package model

type SCTInfo struct {
	Sid   int     `json:"sid"`
	Cid   int     `json:"cid"`
	Tid   int     `json:"tid"`
	Sname string  `json:"sname"`
	Tname string  `json:"tname"`
	Cname string  `json:"cname"`
	Grade float64 `json:"grade"`
	Term  string  `json:"term"`
}
