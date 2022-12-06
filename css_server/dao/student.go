package dao

import (
	"CSS/model"
	"fmt"
)

func FindAllStu() ([]model.Student, error) {
	var stus []model.Student
	dbRes := db.Model(&model.Student{}).Select("*").Find(&stus)
	err := dbRes.Error
	if err != nil {
		return stus, err
	}
	return stus, nil
}

func FindStuById(sid int) (model.Student, error) {
	var stu model.Student
	dbRes := db.Model(&model.Student{}).Select("*").
		Where("sid = ?", sid).Find(&stu)
	err := dbRes.Error
	if err != nil {
		return model.Student{}, err
	}
	return stu, nil
}

func FindStuBySearch(stu model.Student, fuzzy int) ([]model.Student, error) {
	var stus []model.Student
	dbRes := db.Model(&model.Student{}).Select("*")
	if stu.Sid != 0 {
		dbRes = dbRes.Where("sid = ?", stu.Sid)
	}
	if stu.Sname != "" {
		if fuzzy == 0 {
			dbRes = dbRes.Where("sname = ?", stu.Sname)
		} else if fuzzy == 1 {
			dbRes = dbRes.Where("sname like ?", "%"+stu.Sname+"%")
		}
	}
	dbRes = dbRes.Find(&stus)
	err := dbRes.Error
	if err != nil {
		return stus, err
	}
	return stus, nil
}

func UpdateStuById(stu model.Student) error {
	dbRes := db.Model(&model.Student{}).Where("sid = ?", stu.Sid).
		Updates(map[string]interface{}{"sname": stu.Sname, "Password": stu.Password})
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func InsertStu(stu model.Student) error {
	deRes := db.Create(&stu)
	err := deRes.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteStuById(sid int) error {
	dbRes := db.Delete(&model.Student{}, sid)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}
