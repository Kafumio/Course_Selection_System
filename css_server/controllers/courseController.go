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

var cou model.Course

func FindCouBySearch(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var Map map[string]interface{}
	json.Unmarshal(body, &Map)
	couList, err := services.FindCouBySearch(Map)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索课程失败")
		return
	}
	util.RespSuccessfulWithData(ctx, couList)
	return
}

func FindCouById(ctx *gin.Context) {
	cid, _ := strconv.Atoi(ctx.Param("cid"))
	couList, err := services.FindCouById(cid)
	if err != nil {
		util.RespErrorWithData(ctx, "搜索课程 by cid失败")
		return
	}
	util.RespSuccessfulWithData(ctx, couList)
	return
}

func AddCourse(ctx *gin.Context) {
	ctx.BindJSON(&cou)
	if services.InsertCou(cou) == false {
		util.RespErrorWithData(ctx, "插入课程失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func DeleteCouById(ctx *gin.Context) {
	cid, _ := strconv.Atoi(ctx.Param("cid"))
	if services.DeleteCouById(cid) == false {
		util.RespErrorWithData(ctx, "删除课程失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}

func UpdateCou(ctx *gin.Context) {
	ctx.BindJSON(&cou)
	fmt.Println("更新教师 ", cou)
	if services.UpdateCouById(cou) == false {
		util.RespErrorWithData(ctx, "更新课程失败")
		return
	}
	util.RespSuccessful(ctx)
	return
}
