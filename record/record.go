package record

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RecordItem struct {
	Timestamp float64   `form:"timestamp" json:"timestamp" binding:"required"`
	Vectors   []float64 `form:"vectors" json:"vectors" binding:"required"`
}

type Record struct {
	Datas []RecordItem `form:"datas" json:"datas" binding:"required"`
}

type HexagramItem struct {
	Hexagram string `form:"hexagram" json:"hexagram" binding:"required"`
	Value    int    `form:"value" json:"value" binding:"required"`
}

var counter = 0
var database = map[int]Record{}
var tables = map[string][]float64{
	"1": {0.5, 0.5, 0.5, 0.5, 0.5, 0.5},
	"2": {0.5, 0.5, 50, 50, 50, 50},
}

var hexagram = [][]string{
	{"乾", "履", "同人", "無妄", "姤", "訟", "遁", "否"},
	{"夬", "兑", "革", "隨", "大過", "困", "咸", "萃"},
	{"大有", "睽", "離", "噬嗑", "鼎", "未濟", "旅", "晉"},
	{"大壯", "歸妹", "豐", "震", "恆", "解", "小過", "豫"},
	{"小畜", "中孚", "家人", "益", "巽", "渙", "漸", "觀"},
	{"需", "節", "既濟", "屯", "井", "坎", "蹇", "比"},
	{"大畜", "損", "賁", "頤", "蠱", "蒙", "艮", "剝"},
	{"泰", "臨", "明夷", "復", "升", "師", "謙", "坤"},
}

func GetRecord(context *gin.Context) {
	fmt.Println("log")

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

func GetRecordSequence(context *gin.Context) {
	id, key, ok := checkIdAndTable(context)
	if !ok {
		return
	}

	record, ok := database[id]
	table, ok := tables[key]
	var array []string
	for _, item := range record.Datas {
		above, below := mapHexagram(table, item.Vectors)
		array = append(array, hexagram[above][below])
	}

	context.JSON(http.StatusOK, gin.H{
		"datas": array,
	})
}

func GetRecordFrequency(context *gin.Context) {
	id, key, ok := checkIdAndTable(context)
	if !ok {
		return
	}

	record, ok := database[id]
	table, ok := tables[key]
	datas := getSurface3dChartData(table, record)

	var array []HexagramItem
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if datas[i][j] == 0 {
				continue
			}
			item := HexagramItem{
				Hexagram: hexagram[i][j],
				Value:    datas[i][j],
			}
			array = append(array, item)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"datas": array,
	})
}

func checkIdAndTable(context *gin.Context) (int, string, bool) {
	id, ok := queryInt("id", context)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, "", false
	}

	if _, ok := database[id]; !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, "", false
	}

	key := context.Query("table")
	if key == "" {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, "", false
	}

	if _, ok := tables[key]; !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, "", false
	}

	return id, key, true
}

func AddRecord(context *gin.Context) {
	var json Record
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for _, item := range json.Datas {
		if len(item.Vectors) != 6 {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	counter += 1
	database[counter] = json

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      counter,
	})
}
