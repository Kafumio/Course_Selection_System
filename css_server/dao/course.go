package dao

import (
	"CSS/model"
	"fmt"
)

func FindCouBySearch(cid, fuzzy, lowBound, highBound int, cname string) ([]model.Course, error) {
	var cous []model.Course
	dbRes := db.Model(&model.Course{}).Select("*")
	if cid != 0 {
		dbRes = dbRes.Where("cid = ?", cid)
	}
	if cname != "" {
		if fuzzy == 0 {
			dbRes = dbRes.Where("cname = ?", cname)
		} else if fuzzy == 1 {
			dbRes = dbRes.Where("cname like ?", "%"+cname+"%")
		}
	}
	if lowBound > 0 {
		dbRes = dbRes.Where("ccredit >= ?", lowBound)
	}
	if highBound > 0 {
		dbRes = dbRes.Where("ccredit <= ?", highBound)
	}
	dbRes = dbRes.Find(&cous)
	err := dbRes.Error
	if err != nil {
		return nil, err
	}
	return cous, nil
}

func InsertCou(cou model.Course) error {
	deRes := db.Create(&cou)
	err := deRes.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateCouById(cou model.Course) error {
	dbRes := db.Model(&model.Student{}).Where("cid = ?", cou.Cid).
		Updates(map[string]interface{}{"cname": cou.Cname, "ccredit": cou.Ccredit})
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteCouById(cid int) error {
	dbRes := db.Delete(&model.Course{}, cid)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}
