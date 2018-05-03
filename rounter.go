package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", getPing)
		v1.GET("/record", getRecord)
		v1.POST("/record", addRecord)
		v1.GET("/chart3d", getChart3d)
	}
}

func getPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
