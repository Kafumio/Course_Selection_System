package controllers

import (
	"CSS/model"
	"CSS/services"
	"CSS/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

var ct model.CourseTeacher

func InsertCT(ctx *gin.Context) {
	cid, _ := strconv.Atoi(ctx.Param("cid"))
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	term := ctx.Param("term")
	if services.InsertCourseTeacher(cid, tid, term) == false {
		util.RespErrorWithData(ctx, "插入课程-教师关系失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func FindMyCourse(ctx *gin.Context) {
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	term := ctx.Param("term")
	fmt.Println("查询教师课程：", tid, " ", term)
	ctList, err := services.FindMyCourse(tid, term)
	if err != nil {
		util.RespErrorWithData(ctx, "查询教师课程失败")
		return
	}
	util.RespSuccessfulWithData(ctx, ctList)
	return
}

func FindCTById(ctx *gin.Context) {
	cid, _ := strconv.Atoi(ctx.Param("cid"))
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	term := ctx.Param("term")
	fmt.Println("查询课程-教师关系：", cid, " ", tid, " ", term)
	ct, err := services.FindCTById(cid, tid, term)
	if err != nil {
		util.RespErrorWithData(ctx, "查询教师课程失败")
		return
	}
	fmt.Println(ct)
	util.RespSuccessfulWithData(ctx, ct)
	return
}

func FindCTInfo(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var Map map[string]interface{}
	json.Unmarshal(body, &Map)
	ctList, err := services.FindCourseTeacherInfo(Map)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索课程-教师关系信息失败")
		return
	}
	util.RespSuccessfulWithData(ctx, ctList)
	return
}

func DeleteCTById(ctx *gin.Context) {
	ctx.BindJSON(&ct)
	if services.DeleteCTById(ct) == false {
		util.RespErrorWithData(ctx, "删除课程-教师关系失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}
