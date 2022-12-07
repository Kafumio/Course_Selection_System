package services

import (
	"CSS/dao"
	"CSS/model"
	"fmt"
)

func FindCTInfoByStuId(sid int, term string) ([]model.CourseTeacherInfo, error) {
	return dao.FindCTInfoByStuId(sid, term)
}

func FindAllTerm() ([]string, error) { return dao.FindAllTerm() }

func IsSCTExist(sct model.StudentCourseTeacher) bool {
	sctList, err := dao.FindBySCT(sct)
	if err != nil {
		return false
	}
	return len(sctList) != 0
}

func InsertSCT(sct model.StudentCourseTeacher) bool {
	return dao.InsertSCT(sct) == nil
}

func DeleteBySCT(sct model.StudentCourseTeacher) bool {
	return dao.DeleteSCT(sct) == nil
}

func DeleteSCTById(sid, cid, tid int, term string) bool {
	sct := model.StudentCourseTeacher{Sid: sid, Cid: cid, Tid: tid, Term: term}
	return dao.InsertSCT(sct) == nil
}

func FindByIdWithTerm(sid, cid, tid int, term string) (model.SCTInfo, error) {
	sct, err := dao.FindSCTBySearch(sid, 0, cid, 0, tid, 0,
		0, 0, "", "", "", "")
	return sct[0], err
}

func UpdateSCTById(sid, cid, tid, grade int, term string) bool {
	return dao.UpdateSCTGradeById(sid, cid, tid, grade, term) == nil
}

func FindSCTBySearch(Map map[string]interface{}) ([]model.SCTInfo, error) {
	sid, tid, cid, sFuzzy, tFuzzy, cFuzzy := 0, 0, 0, 0, 0, 0
	lowBound, highBound := 0, 0
	sname, tname, cname, term := "", "", "", ""

	if _, ok := Map["sid"]; ok {
		tmp, ok := Map["sid"].(float64)
		if ok {
			sid = int(tmp)
		}
	}
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
	if _, ok := Map["sname"]; ok {
		sname, ok = Map["sname"].(string)
	}
	if _, ok := Map["tname"]; ok {
		tname, ok = Map["tname"].(string)
	}
	if _, ok := Map["cname"]; ok {
		cname, ok = Map["cname"].(string)
	}
	if _, ok := Map["sFuzzy"]; ok {
		if Map["sFuzzy"] == true {
			sFuzzy = 1
		} else {
			sFuzzy = 0
		}
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
	if _, ok := Map["term"]; ok {
		term, ok = Map["term"].(string)
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
	fmt.Println("SCT查询：", Map)
	return dao.FindSCTBySearch(sid, sFuzzy, cid, cFuzzy, tid, tFuzzy,
		lowBound, highBound, sname, cname, tname, term)
}
