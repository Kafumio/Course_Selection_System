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

var tea model.Teacher

func AddTeacher(ctx *gin.Context) {
	ctx.ShouldBind(&tea)
	if services.InsertTea(tea) == false {
		util.RespErrorWithData(ctx, "插入教师失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func LoginTea(ctx *gin.Context) {
	ctx.ShouldBind(&tea)
	fmt.Println("正在验证教师登陆 ", tea)
	t, err := services.FindTeaById(tea.Tid)
	if err != nil {
		util.RespErrorWithData(ctx, "登陆失败")
		return
	}
	if t == (model.Teacher{}) || t.Password != tea.Password {
		util.RespErrorWithData(ctx, "登陆类型或id密码错误")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func FindTeaById(ctx *gin.Context) {
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	fmt.Println("正在查询教师信息 by id：", tid)
	t, err := services.FindTeaById(tid)
	if err != nil {
		util.RespErrorWithData(ctx, "查询失败")
		return
	}
	util.RespSuccessfulWithData(ctx, t)
}

func FindTeaBySearch(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var Map map[string]interface{}
	json.Unmarshal(body, &Map)
	teaList, err := services.FindTeaBySearch(Map)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索教师失败")
		return
	}
	util.RespSuccessfulWithData(ctx, teaList)
}

func DeleteTeaById(ctx *gin.Context) {
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	fmt.Println("正在删除教师 tid：", tid)
	if services.DeleteTeaById(tid) == false {
		util.RespErrorWithData(ctx, "删除教师失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func UpdateTea(ctx *gin.Context) {
	ctx.ShouldBind(&tea)
	fmt.Println("更新教师 ", tea)
	if services.UpdateTeaById(tea) == false {
		util.RespErrorWithData(ctx, "更新教师失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}
