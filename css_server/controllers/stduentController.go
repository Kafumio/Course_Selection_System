package controllers

import (
	"CSS/model"
	"CSS/services"
	"CSS/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var stu model.Student

func AddStudent(ctx *gin.Context) {
	ctx.ShouldBind(&stu)
	fmt.Println("正在添加学生对象 ", stu)
	if services.InsertStu(stu) == false {
		util.RespErrorWithData(ctx, "添加学生失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func LoginStu(ctx *gin.Context) {
	ctx.ShouldBind(&stu)
	fmt.Println("正在验证学生登陆 ", stu)
	s, err := services.FindStuById(stu.Sid)
	if err != nil {
		util.RespErrorWithData(ctx, "登陆失败")
		return
	}
	if s == (model.Student{}) || s.Password != stu.Password {
		util.RespErrorWithData(ctx, "登陆类型或id密码错误")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func FindStuBySearch(ctx *gin.Context) {
	ctx.ShouldBind(&stu)
	var fuzzy int
	if stu == (model.Student{}) {
		fuzzy = 0
	} else {
		fuzzy = 1
	}
	stuList, err := services.FindStuBySearch(fuzzy, stu.Sid, stu.Sname)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索失败")
		return
	}
	util.RespSuccessfulWithData(ctx, stuList)
	return
}

func FindStuById(ctx *gin.Context) {
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	fmt.Println("正在查询学生信息 by id：", sid)
	s, err := services.FindStuById(sid)
	if err != nil {
		util.RespErrorWithData(ctx, "查询失败")
		return
	}
	util.RespSuccessfulWithData(ctx, s)
}

func FindStuByPage(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Param("page"))
	size, _ := strconv.Atoi(ctx.Param("size"))
	fmt.Println("查询学生列表分页", page, "", size)
	stuList, err := services.FindStuByPage(page, size)
	if err != nil {
		util.RespErrorWithData(ctx, "查询学生列表分页失败")
		return
	}
	util.RespSuccessfulWithData(ctx, stuList)
	return
}

func GetStuLength(ctx *gin.Context) {
	util.RespSuccessfulWithData(ctx, services.GetStuLength())
	return
}

func DeleteStuById(ctx *gin.Context) {
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	fmt.Println("正在删除学生 sid：", sid)
	if services.DeleteStuById(sid) == false {
		util.RespErrorWithData(ctx, "删除学生失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func UpdateStu(ctx *gin.Context) {
	ctx.ShouldBind(&stu)
	fmt.Println("更新学生信息：", stu)
	if services.UpdateStuById(stu) == false {
		util.RespErrorWithData(ctx, "更新学生信息失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}
