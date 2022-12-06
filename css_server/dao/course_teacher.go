package dao

import (
	"CSS/model"
	"fmt"
)

func InsertCourseTeacher(cid, tid int, term string) error {
	ct := model.CourseTeacher{Cid: cid, Tid: tid, Term: term}
	dbRes := db.Create(&ct)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func FindCTBySearch(cid, tid int, term string) ([]model.CourseTeacher, error) {
	cts := []model.CourseTeacher{}
	dbRes := db.Model(&model.CourseTeacher{}).Select("*")
	if cid != 0 {
		dbRes = dbRes.Where("cid = ?", cid)
	}
	if tid != 0 {
		dbRes = dbRes.Where("tid = ?", tid)
	}
	if term != "" {
		dbRes = dbRes.Where("term = ?", term)
	}
	dbRes = dbRes.Find(&cts)
	err := dbRes.Error
	if err != nil {
		return cts, err
	}
	return cts, nil
}

func FindMyCourse(tid int, term string) ([]model.Course, error) {
	cous := []model.Course{}
	dbRes := db.Raw("select c.cid, c.cname, t.tid, tname, ccredit from (c,t) inner "+
		"join ct on c.cid = ct.cid and t.tid = ct.tid where ct.tid = ? and "+
		"ct.term = ?", tid, term).Scan(&cous)
	err := dbRes.Error
	if err != nil {
		return cous, err
	}
	return cous, nil
}

func FindCTById(cid int, tid int, term string) (model.CourseTeacher, error) {
	ct := model.CourseTeacher{}
	dbRes := db.Model(&model.CourseTeacher{}).Select("*").Where("cid = ? and "+
		"tid = ? and term = ?", cid, tid, term).Find(&ct)
	err := dbRes.Error
	if err != nil {
		return ct, err
	}
	return ct, nil
}

func FindCourseTeacherInfo(tid, tFuzzy, cid, cFuzzy int, tname, cname string) ([]model.CourseTeacherInfo, error) {
	ctInfos := []model.CourseTeacherInfo{}
	/*
		dbRes := db.Table("ct").Select("c.cid, cname, t.tid, tname, c.ccredit").
			Joins("inner join c on c.cid = ct.cid").
			Joins("inner join t on t.tid = ct.tid")
	*/
	dbRes := db.Raw("select * from ct_info")
	if tid != 0 {
		dbRes.Where("t.tid = ?", tid)
	}
	if tname != "" {
		if tFuzzy == 0 {
			dbRes = dbRes.Where("t.tname = ?", tname)
		} else if tFuzzy == 1 {
			dbRes = dbRes.Where("t.tname like ?", "%"+tname+"%")
		}
	}
	if cid != 0 {
		dbRes.Where("c.cid = ?", cid)
	}
	if cname != "" {
		if cFuzzy == 0 {
			dbRes = dbRes.Where("c.cname = ?", cname)
		} else if cFuzzy == 1 {
			dbRes = dbRes.Where("c.cname like ?", "%"+cname+"%")
		}
	}
	dbRes.Find(&ctInfos)
	err := dbRes.Error
	if err != nil {
		return ctInfos, err
	}
	return ctInfos, nil
}

func DeleteCTById(ct model.CourseTeacher) error {
	dbRes := db.Raw("delete from ct where cid = ? and tid = ?", ct.Cid, ct.Tid)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}
