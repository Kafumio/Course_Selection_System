package dao

import (
	"CSS/model"
	"fmt"
)

func FindAllTea() ([]model.Teacher, error) {
	var teas []model.Teacher
	dbRes := db.Model(&model.Teacher{}).Select("*").Find(&teas)
	err := dbRes.Error
	if err != nil {
		return teas, err
	}
	return teas, nil
}

func FindTeaById(tid int) (model.Teacher, error) {
	var tea model.Teacher
	dbRes := db.Model(&model.Teacher{}).Select("*").Where("tid = ?", tid).Find(&tea)
	err := dbRes.Error
	if err != nil {
		return tea, err
	}
	return tea, nil
}

func FindTeaBySearch(tid, fuzzy int, tname string) ([]model.Teacher, error) {
	var teas []model.Teacher
	dbRes := db.Model(&model.Teacher{}).Select("*")
	if tid != 0 {
		dbRes = dbRes.Where("tid = ?", tid)
	}
	if tname != "" {
		if fuzzy == 0 {
			dbRes = dbRes.Where("tname = ?", tname)
		} else if fuzzy == 1 {
			dbRes = dbRes.Where("tname like ?", "%"+tname+"%")
		}
	}
	dbRes = dbRes.Find(&teas)
	err := dbRes.Error
	if err != nil {
		return teas, err
	}
	return teas, nil
}

func UpdateTeaById(tea model.Teacher) error {
	dbRes := db.Model(&model.Teacher{}).Where("tid = ?", tea.Tid).
		Updates(map[string]interface{}{"tname": tea.Tname, "Password": tea.Password})
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func InsertTea(tea model.Teacher) error {
	deRes := db.Create(&tea)
	err := deRes.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteTeaById(tid int) error {
	dbRes := db.Delete(&model.Teacher{}, tid)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}
