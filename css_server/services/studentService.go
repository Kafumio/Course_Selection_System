package services

import (
	"CSS/dao"
	"CSS/model"
)

func FindStuByPage(num, size int) ([]model.Student, error) {
	// num：第几页，size：一页多大
	// num：从零开始
	curPage := []model.Student{}
	studentList, err := dao.FindAllStu()
	if err != nil {
		return nil, err
	}
	start := size * num
	end := size * (num + 1)
	listLen := len(studentList)
	for i := start; i < end && i < listLen; i++ {
		curPage = append(curPage, studentList[i])
	}
	return curPage, err
}

func FindStuBySearch(fuzzy, sid int, sname string) ([]model.Student, error) {
	stu := model.Student{Sid: sid, Sname: sname}
	return dao.FindStuBySearch(stu, fuzzy)
}

func GetStuLength() int {
	stuList, err := dao.FindAllStu()
	if err != nil {
		return 0
	}
	return len(stuList)
}

func FindStuById(sid int) (model.Student, error) {
	return dao.FindStuById(sid)
}

func UpdateStuById(stu model.Student) bool {
	err := dao.UpdateStuById(stu)
	return err == nil
}

func InsertStu(stu model.Student) bool {
	err := dao.InsertStu(stu)
	return err == nil
}

func DeleteStuById(sid int) bool {
	err := dao.DeleteStuById(sid)
	return err == nil
}
