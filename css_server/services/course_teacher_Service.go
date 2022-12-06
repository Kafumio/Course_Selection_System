package services

import (
	"CSS/dao"
	"CSS/model"
	"fmt"
	"strconv"
)

func InsertCourseTeacher(cid, tid int, term string) bool {
	return dao.InsertCourseTeacher(cid, tid, term) == nil
}

func FindMyCourse(tid int, term string) ([]model.Course, error) {
	return dao.FindMyCourse(tid, term)
}

func FindCTById(cid, tid int, term string) (model.CourseTeacher, error) {
	return dao.FindCTById(cid, tid, term)
}

func FindCourseTeacherInfo(Map map[string]interface{}) ([]model.CourseTeacherInfo, error) {
	tid, cid, tFuzzy, cFuzzy := 0, 0, 0, 0
	tname, cname := "", ""

	if _, ok := Map["tid"]; ok {
		tmp, ok := Map["tid"].(float64)
		if ok {
			tid = int(tmp)
		}
	}
	if _, ok := Map["cid"]; ok {
		tmp, ok := Map["cid"].(float64)
		if ok {
			cid = int(tmp)
		}
	}
	if _, ok := Map["tname"]; ok {
		tname, ok = Map["tname"].(string)
	}
	if _, ok := Map["cname"]; ok {
		cname, ok = Map["cname"].(string)
	}
	if _, ok := Map["tFuzzy"]; ok {
		if Map["tFuzzy"] == true {
			tFuzzy = 1
		} else {
			tFuzzy = 0
		}
	}
	if _, ok := Map["cFuzzy"]; ok {
		if Map["cFuzzy"] == true {
			cFuzzy = 1
		} else {
			cFuzzy = 0
		}
	}
	fmt.Println("ct 模糊查询", Map)
	fmt.Printf("%d %d %d %d %s %s\n", tid, tFuzzy, cid, cFuzzy, tname, cname)
	return dao.FindCourseTeacherInfo(tid, tFuzzy, cid, cFuzzy, tname, cname)
}

func FindBySearch(cid, tid int, term string) ([]model.CourseTeacher, error) {
	return dao.FindCTBySearch(cid, tid, term)
}

func FindBySearchMap(Map map[string]string) ([]model.CourseTeacher, error) {
	tid, cid := 0, 0
	term := ""
	if _, ok := Map["tid"]; ok {
		tid, _ = strconv.Atoi(Map["tid"])
	}
	if _, ok := Map["cid"]; ok {
		cid, _ = strconv.Atoi(Map["cid"])
	}
	if _, ok := Map["term"]; ok {
		term = Map["term"]
	}
	fmt.Println("开课表查询：", Map)
	return dao.FindCTBySearch(cid, tid, term)
}

func DeleteCTById(ct model.CourseTeacher) bool {
	return dao.DeleteCTById(ct) == nil
}
