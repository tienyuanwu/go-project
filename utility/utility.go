package utility

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryInt(key string, context *gin.Context) (int64, bool) {
	str := context.Query(key)
	if str == "" {
		return -1, false
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		return -1, false
	}

	return int64(value), true
}
