package main

import (
	"./record"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", getPing)
		v1.GET("/record", record.GetRecord)
		v1.POST("/record", record.AddRecord)
		v1.GET("/chart3d", record.GetChart3d)
		v1.GET("/chart/frequency", record.GetChartFrequncy)
		v1.GET("/chart/sequnce", record.GetChartSequnce)
		v1.GET("/record/frequency", record.GetRecordFrequency)
		v1.GET("/record/sequnce", record.GetRecordSequence)
	}
}

func getPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
