package dao

import (
	"CSS/model"
	"fmt"
	"gorm.io/gorm"
)

func FindCTInfoByStuId(sid int, term string) ([]model.CourseTeacherInfo, error) {
	ctInfos := []model.CourseTeacherInfo{}
	/*
		dbRes := db.Raw("select c.cid, c.cname, t.tid, t.tname, sct.grade, c.ccredit from"+
			"sct inner join t on sct.tid = t.tid "+
			"inner join c on sct.cid = c.cid where "+
			"sct.sid = ? and sct.term = ?", sid, term).Scan(&ctInfos)
	*/
	dbRes := db.Table("sct").Select("c.cid, cname, t.tid, tname, sct.grade, c.ccredit").
		Joins("inner join c on sct.cid = c.cid").
		Joins("inner join t on sct.tid = t.tid").
		Where("sct.sid = ? and sct.term = ?", sid, term).Find(&ctInfos)
	err := dbRes.Error
	if err != nil {
		return nil, err
	}
	return ctInfos, err
}

func FindSCTBySearch(sid, sFuzzy, cid, cFuzzy, tid, tFuzzy,
	lowBound, highBound int, sname, cname, tname, term string) ([]model.SCTInfo, error) {
	sctInfos := []model.SCTInfo{}
	var dbRes *gorm.DB
	if sid == 0 && tid == 0 && cid == 0 && sname == "" && tname == "" && cname == "" {
		dbRes = db.Raw("select * from sct_info")
	} else {
		dbRes = db.Table("sct").Select("s.sid, sname, c.cid, " +
			"cname, t.tid, tname, sct.grade, sct.term").
			Joins("inner join s on sct.sid = s.sid").
			Joins("inner join c on sct.cid = c.cid").
			Joins("inner join t on sct.tid = t.tid")
	}
	if sid != 0 {
		dbRes = dbRes.Where("s.sid = ?", sid)
	}
	if cid != 0 {
		dbRes = dbRes.Where("c.cid = ?", cid)
	}
	if tid != 0 {
		dbRes = dbRes.Where("t.tid = ?", tid)
	}
	if sname != "" {
		if sFuzzy == 0 {
			dbRes = dbRes.Where("s.sname = ?", sname)
		} else if sFuzzy == 1 {
			dbRes = dbRes.Where("s.sname like ?", "%"+sname+"%")
		}
	}
	if cname != "" {
		if cFuzzy == 0 {
			dbRes = dbRes.Where("c.cname = ?", cname)
		} else if cFuzzy == 1 {
			dbRes = dbRes.Where("c.cname like ?", "%"+cname+"%")
		}
	}
	if tname != "" {
		if tFuzzy == 0 {
			dbRes = dbRes.Where("t.tname = ?", tname)
		} else if tFuzzy == 1 {
			dbRes = dbRes.Where("t.tname like ?", "%"+tname+"%")
		}
	}
	if term != "" {
		dbRes = dbRes.Where("sct.term = ?", term)
	}
	if lowBound != 0 {
		dbRes = dbRes.Where("sct.grade >= ?", lowBound)
	}
	if highBound != 0 {
		dbRes = dbRes.Where("sct.grade <= ?", highBound)
	}
	dbRes.Find(&sctInfos)
	err := dbRes.Error
	if err != nil {
		return sctInfos, err
	}
	return sctInfos, nil
}

func FindAllTerm() ([]string, error) {
	terms := []string{}
	dbRes := db.Raw("select distinct term from sct").Scan(&terms)
	err := dbRes.Error
	if err != nil {
		return nil, err
	}
	return terms, err
}

func FindBySCT(sct model.StudentCourseTeacher) ([]model.StudentCourseTeacher, error) {
	scts := []model.StudentCourseTeacher{}
	dbRes := db.Model(&model.StudentCourseTeacher{}).Select("*").Where("sid = ? "+
		"and cid = ? and tid = ? and term = ?", sct.Sid, sct.Cid, sct.Tid, sct.Term).Find(&scts)
	err := dbRes.Error
	if err != nil {
		return nil, err
	}
	return scts, err
}

func InsertSCT(sct model.StudentCourseTeacher) error {
	err := db.Create(&sct).Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateSCTGradeById(sid, cid, tid, grade int, term string) error {
	err := db.Model(&model.StudentCourseTeacher{}).Where("sid = ? and cid = ? and tid = ? "+
		"and term = ?", sid, cid, tid, term).Update("grade", grade).Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteSCT(sct model.StudentCourseTeacher) error {
	err := db.Delete(&model.StudentCourseTeacher{}, sct).Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}
