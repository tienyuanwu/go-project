package record

import (
	"../table"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Color struct {
	R int `form:"r" json:"r" binding:"required"`
	G int `form:"g" json:"g" binding:"required"`
	B int `form:"b" json:"b" binding:"required"`
}

type ChartItem struct {
	X     int   `form:"x" json:"x" binding:"required"`
	Y     int   `form:"y" json:"y" binding:"required"`
	Z     int   `form:"z" json:"z" binding:"required"`
	Color Color `form:"color" json:"color" binding:"required"`
}

var defaultColors = []Color{
	Color{255, 0, 0},
	Color{255, 165, 0},
	Color{255, 255, 0},
	Color{0, 255, 255},
	Color{0, 0, 255},
	Color{43, 0, 255},
	Color{87, 0, 255},
}

func GetChartFrequncy(context *gin.Context) {
	id, key, ok := checkIdAndTable(context)
	if !ok {
		return
	}

	record, ok := database[id]
	table, ok := table.GetTable(key)
	result := getSurface3dChartData(record, table)

	datas := []ChartItem{}
	colorsLength := len(defaultColors)
	colorIndex := 0
	for i, array := range result {
		for j, value := range array {
			if value == 0 {
				continue
			}
			color := defaultColors[colorIndex%colorsLength]
			item := ChartItem{i, j, value, color}
			datas = append(datas, item)
			colorIndex += 1
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"datas": datas,
	})
}

func GetChartSequence(context *gin.Context) {
	id, key, ok := checkIdAndTable(context)
	if !ok {
		return
	}

	record, ok := database[id]
	table, ok := table.GetTable(key)

	datas := []ChartItem{}
	colorsLength := len(defaultColors)
	for i, item := range record.Datas {
		above, below := mapHexagram(table, item.Vectors)
		color := defaultColors[i%colorsLength]
		item := ChartItem{above, below, i, color}
		datas = append(datas, item)
	}

	context.JSON(http.StatusOK, gin.H{
		"datas": datas,
	})
}

func GetChart3d(context *gin.Context) {
	id, key, ok := checkIdAndTable(context)
	if !ok {
		return
	}

	record, ok := database[id]
	table, ok := table.GetTable(key)
	datas := getSurface3dChartData(record, table)

	context.JSON(http.StatusOK, gin.H{
		"datas": datas,
	})
}

func getSurface3dChartData(record Record, table [6]float64) [8][8]int {
	datas := [8][8]int{
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
		above, below := mapHexagram(table, item.Vectors)
		datas[above][below] += 1
	}

	return datas
}
