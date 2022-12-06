package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithData(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info":  false,
		"error": err,
	})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": true,
	})
}

func RespSuccessfulWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info":    true,
		"content": data,
	})
}