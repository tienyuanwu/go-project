package suggestion

import (
	"github.com/gin-gonic/gin"
	//	"net/http"
)

type Suggestion struct {
	Name string            `form:"name" json:"name" binding:"required"`
	Data map[string]string `form:"datas" json:"datas" binding:"required"`
}

var database = map[string]Suggestion{}

func Add(context *gin.Context) {
}

func Get(context *gin.Context) {
}
