package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func queryInt(key string, context *gin.Context) (int, bool) {
	str := context.Query(key)
	if str == "" {
		return -1, false
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		return -1, false
	}

	return value, true
}
