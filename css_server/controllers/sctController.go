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

var sct model.StudentCourseTeacher

func InsertSCT(ctx *gin.Context) {
	ctx.ShouldBind(&sct)
	if services.IsSCTExist(sct) {
		util.RespErrorWithData(ctx, "禁止重复选课")
		return
	}
	fmt.Println("正在保存sct记录：", sct)
	if services.InsertSCT(sct) == false {
		util.RespErrorWithData(ctx, "选课失败，请联系管理员")
		return
	}
	util.RespSuccessfulWithData(ctx, "选课成功")
	return
}

func FindCTInfoBySid(ctx *gin.Context) {
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	term := ctx.Param("term")
	sctList, err := services.FindCTInfoByStuId(sid, term)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索sct by id失败")
		return
	}
	util.RespSuccessfulWithData(ctx, sctList)
	return
}

func FindAllTerm(ctx *gin.Context) {
	terms, _ := services.FindAllTerm()
	util.RespSuccessfulWithData(ctx, terms)
	return
}

func DeleteBySCT(ctx *gin.Context) {
	ctx.ShouldBind(&sct)
	fmt.Println("正在删除sct记录：", sct)
	if services.DeleteBySCT(sct) == false {
		util.RespErrorWithData(ctx, "删除sct记录失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func FindSCTBySearch(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var Map map[string]interface{}
	json.Unmarshal(body, &Map)
	sctList, err := services.FindSCTBySearch(Map)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索sct记录失败")
		return
	}
	util.RespSuccessfulWithData(ctx, sctList)
	return
}

func FindSCTById(ctx *gin.Context) {
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	cid, _ := strconv.Atoi(ctx.Param("cid"))
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	term := ctx.Param("term")
	sctInfo, err := services.FindByIdWithTerm(sid, cid, tid, term)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索sct记录 by id失败")
		return
	}
	util.RespSuccessfulWithData(ctx, sctInfo)
	return
}

func UpdateSCTById(ctx *gin.Context) {
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	cid, _ := strconv.Atoi(ctx.Param("cid"))
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	term := ctx.Param("term")
	grade, _ := strconv.Atoi(ctx.Param("grade"))
	if services.UpdateSCTById(sid, cid, tid, grade, term) == false {
		util.RespErrorWithData(ctx, "更新sct记录失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func DeleteSCTById(ctx *gin.Context) {
	sid, _ := strconv.Atoi(ctx.Param("sid"))
	cid, _ := strconv.Atoi(ctx.Param("cid"))
	tid, _ := strconv.Atoi(ctx.Param("tid"))
	term := ctx.Param("term")
	if services.DeleteSCTById(sid, cid, tid, term) == false {
		util.RespErrorWithData(ctx, "删除sct记录 by id失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}
