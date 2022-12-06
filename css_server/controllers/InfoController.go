package controllers

import (
	"CSS/util"
	"github.com/gin-gonic/gin"
)

const CURRENT_TERM = "22-秋季学期"
const FORBID_COURSE_SELECTION = false

func GetCurrentTerm(c *gin.Context) {
	util.RespSuccessfulWithData(c, CURRENT_TERM)
}

func GetForbidCourseSelection(c *gin.Context) {
	util.RespSuccessfulWithData(c, FORBID_COURSE_SELECTION)
}
