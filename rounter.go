package main

import (
	"./record"
	"./table"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", getPing)
		v1.GET("/record", record.Get)
		v1.POST("/record", record.Add)
		v1.GET("/chart3d", record.GetChart3d)
		v1.GET("/chart/frequency", record.GetChartFrequncy)
		v1.GET("/chart/sequence", record.GetChartSequence)
		v1.GET("/record/frequency", record.GetRecordFrequency)
		v1.GET("/record/sequnce", record.GetRecordSequence)
		v1.GET("/table", table.Get)
		v1.POST("/table", table.Add)
	}
}

func getPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
