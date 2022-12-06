package services

import (
	"CSS/dao"
	"CSS/model"
	"fmt"
)

func FindCouBySearch(Map map[string]interface{}) ([]model.Course, error) {
	cid, lowBound, highBound, fuzzy := 0, 0, 0, 0
	cname := ""

	if _, ok := Map["cid"]; ok {
		tmp, ok := Map["cid"].(float64)
		if ok {
			cid = int(tmp)
		}
	}
	if _, ok := Map["lowBound"]; ok {
		tmp, ok := Map["lowBound"].(float64)
		if ok {
			lowBound = int(tmp)
		}
	}
	if _, ok := Map["highBound"]; ok {
		tmp, ok := Map["highBound"].(float64)
		if ok {
			highBound = int(tmp)
		}
	}
	if _, ok := Map["cname"]; ok {
		cname, ok = Map["cname"].(string)
	}
	if _, ok := Map["fuzzy"]; ok {
		if Map["fuzzy"] == "true" {
			fuzzy = 1
		} else {
			fuzzy = 0
		}
	}
	fmt.Println("查询课程 map：", Map)
	fmt.Printf("%d %s %d %d %d\n", cid, cname, fuzzy, lowBound, highBound)
	return dao.FindCouBySearch(cid, fuzzy, lowBound, highBound, cname)
}

func FindCouById(cid int) ([]model.Course, error) {
	Map := make(map[string]interface{})
	if cid != 0 {
		Map["cid"] = cid
	}
	return FindCouBySearch(Map)
}

func UpdateCouById(cou model.Course) bool {
	err := dao.UpdateCouById(cou)
	return err == nil
}

func InsertCou(cou model.Course) bool {
	err := dao.InsertCou(cou)
	return err == nil
}

func DeleteCouById(cid int) bool {
	err := dao.DeleteCouById(cid)
	return err == nil
}
