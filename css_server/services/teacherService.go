package services

import (
	"CSS/dao"
	"CSS/model"
	"fmt"
)

func FindTeaBySearch(Map map[string]interface{}) ([]model.Teacher, error) {
	tid, fuzzy := 0, 0
	tname := ""
	if _, ok := Map["tid"]; ok {
		tmp, ok := Map["tid"].(float64)
		if ok {
			tid = int(tmp)
		}
	}
	if _, ok := Map["tname"]; ok {
		tname, ok = Map["tname"].(string)
	}
	if _, ok := Map["fuzzy"]; ok {
		if Map["fuzzy"] == "true" {
			fuzzy = 1
		} else {
			fuzzy = 0
		}
	}
	fmt.Println("查询教师 map：", Map)
	fmt.Printf("%d %s %d\n", tid, tname, fuzzy)
	return dao.FindTeaBySearch(tid, fuzzy, tname)
}

func FindTeaByPage(num, size int) ([]model.Teacher, error) {
	// num：第几页，size：一页多大
	// num：从零开始
	curPage := []model.Teacher{}
	tacherList, err := dao.FindAllTea()
	if err != nil {
		return nil, err
	}
	start := size * num
	end := size * (num + 1)
	listLen := len(tacherList)
	for i := start; i < end && i < listLen; i++ {
		curPage = append(curPage, tacherList[i])
	}
	return curPage, err
}

func GetTeaLength() int {
	teaList, err := dao.FindAllTea()
	if err != nil {
		return 0
	}
	return len(teaList)
}

func FindTeaById(tid int) (model.Teacher, error) {
	return dao.FindTeaById(tid)
}

func UpdateTeaById(tea model.Teacher) bool {
	err := dao.UpdateTeaById(tea)
	return err == nil
}

func InsertTea(tea model.Teacher) bool {
	err := dao.InsertTea(tea)
	return err == nil
}

func DeleteTeaById(tid int) bool {
	err := dao.DeleteTeaById(tid)
	return err == nil
}
