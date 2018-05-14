package record

import (
	"../table"
	"../utility"
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

var counter int64 = 0
var database = map[int64]Record{}

func GetFrequencyTable(recordId int64, tableId int64) ([8][8]int, bool) {
	record, ok := database[recordId]
	if !ok {
		return [8][8]int{}, false
	}

	table, ok := table.GetTable(tableId)
	if !ok {
		return [8][8]int{}, false
	}

	return getSurface3dChartData(record, table), true
}

func GetRecordFrequencyArray(recordId int64, tableId int64) ([]HexagramItem, bool) {
	record, ok := database[recordId]
	if !ok {
		return []HexagramItem{}, false
	}

	table, ok := table.GetTable(tableId)
	if !ok {
		return []HexagramItem{}, false
	}

	datas := getSurface3dChartData(record, table)

	var array []HexagramItem
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if datas[i][j] == 0 {
				continue
			}
			item := HexagramItem{
				Hexagram: utility.GetHexagramPostionTable()[i][j],
				Value:    datas[i][j],
			}
			array = append(array, item)
		}
	}

	return array, true
}

func Get(context *gin.Context) {
	keys := make([]int64, len(database))

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
	table, ok := table.GetTable(key)
	var array []string
	for _, item := range record.Datas {
		above, below := mapHexagram(table, item.Vectors)
		array = append(array, utility.GetHexagramPostionTable()[above][below])
	}

	context.JSON(http.StatusOK, gin.H{
		"datas": array,
	})
}

func GetRecordFrequency(context *gin.Context) {
	recordId, ok := utility.QueryInt("record", context)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tableId, ok := utility.QueryInt("table", context)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	array, ok := GetRecordFrequencyArray(recordId, tableId)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"datas": array,
	})
}

func checkIdAndTable(context *gin.Context) (int64, int64, bool) {
	id, ok := utility.QueryInt("record", context)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, -1, false
	}

	if _, ok := database[id]; !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, -1, false
	}

	key, ok := utility.QueryInt("table", context)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, -1, false
	}

	if _, ok := table.GetTable(key); !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return -1, -1, false
	}

	return id, key, true
}

func Add(context *gin.Context) {
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
