package table

import (
	"../utility"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Table struct {
	Id   int64      `form:"id" json:"id" binding:"required"`
	Name string     `form:"name" json:"name" binding:"required"`
	Data [6]float64 `form:"data" json:"data" binding:"required"`
}

var counter int64 = 1
var database = map[int64]Table{
	1: Table{1, "Default", [6]float64{0.5, 0.5, 0.5, 0.5, 0.5, 0.5}},
}

func GetTable(id int64) ([6]float64, bool) {
	table, ok := database[id]
	if ok {
		return table.Data, ok
	} else {
		return [6]float64{0, 0, 0, 0, 0, 0}, false
	}
}

func Add(context *gin.Context) {
	type JsonType struct {
		Name string     `form:"name" json:"name" binding:"required"`
		Data [6]float64 `form:"data" json:"data" binding:"required"`
	}

	var json JsonType
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	counter += 1
	table := Table{counter, json.Name, json.Data}
	database[counter] = table

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      counter,
	})
}

func Get(context *gin.Context) {
	id, ok := utility.QueryInt("id", context)

	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	table, ok := database[id]
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    table,
	})
}

func GetList(context *gin.Context) {
	type JsonType struct {
		Id   int64  `form:"id" json:"id" binding:"required"`
		Name string `form:"name" json:"name" binding:"required"`
	}

	array := []JsonType{}
	for _, value := range database {
		array = append(array, JsonType{value.Id, value.Name})
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    array,
	})
}
