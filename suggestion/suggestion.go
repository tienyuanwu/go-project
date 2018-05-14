package suggestion

import (
	"github.com/gin-gonic/gin"
	//	"net/http"
)

const (
	SuggestionType4  = "t4"
	SuggestionType16 = "t16"
	SuggestionType64 = "t64"
)

type Suggestion struct {
	Name string            `form:"name" json:"name" binding:"required"`
	Type string            `form:"type" json:"type" binding:"required"`
	Data map[string]string `form:"datas" json:"datas" binding:"required"`
}

var counter int64 = 2
var database = map[int64]Suggestion{}

func Add(context *gin.Context) {
}

func Get(context *gin.Context) {
}
