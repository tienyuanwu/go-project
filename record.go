package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecordItem struct {
	Timestamp float64   `form:"timestamp" json:"timestamp" binding:"required"`
	Vectors   []float64 `form:"vectors" json:"vectors" binding:"required"`
	Length    int       `form:"length" json:"length" binding:"required"`
	Table     string    `form:"table" json:"table" binding:"required"`
}

type Record struct {
	Datas []RecordItem `form:"datas" json:"datas" binding:"required"`
}

var counter = 0
var database = map[int]Record{}
var tables = map[string][]float64{
	"1": {0.5, 0.5, 0.5, 0.5, 0.5, 0.5},
}

func getRecord(context *gin.Context) {
	keys := make([]int, len(database))

	i := 0
	for k := range database {
		keys[i] = k
		i++
	}
	context.JSON(http.StatusOK, gin.H{
		"records": keys,
	})
}

func getChart3d(context *gin.Context) {
	id, ok := queryInt("id", context)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	record, ok := database[id]
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	key := context.Query("table")
	if key == "" {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	table, ok := tables[key]
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	datas := getSurface3dChartData(table, record)

	context.JSON(http.StatusOK, gin.H{
		"datas": datas,
	})
}

func getSurface3dChartData(table []float64, record Record) [][]int {
	datas := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	for _, item := range record.Datas {
		above := 0
		below := 0
		for i := 0; i < 6; i++ {
			line := 1
			if item.Vectors[i] > table[i] {
				line = 0
			}

			if i < 3 {
				below |= line << uint(2-i)
			} else {
				above |= line << uint(5-i)
			}
		}
		datas[above][below] += item.Length
	}

	return datas
}

func addRecord(context *gin.Context) {
	var json Record
	if err := context.ShouldBindJSON(&json); err == nil {
		counter += 1
		database[counter] = json

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"id":      counter,
		})

	} else {
		context.AbortWithError(http.StatusBadRequest, err)
	}
}
