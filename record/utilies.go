package record

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
		above, below := mapHexagram(table, item.Vectors)
		datas[above][below] += 1
	}

	return datas
}

func mapHexagram(table []float64, data []float64) (int, int) {
	above := 0
	below := 0
	for i := 0; i < 6; i++ {
		line := 1
		if data[i] > table[i] {
			line = 0
		}

		if i < 3 {
			below |= line << uint(2-i)
		} else {
			above |= line << uint(5-i)
		}
	}
	return above, below
}