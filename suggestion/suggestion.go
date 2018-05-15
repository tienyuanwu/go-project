package suggestion

import (
	"../record"
	"../utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

const (
	SuggestionType4  = "sugg4"
	SuggestionType16 = "sugg16"
	SuggestionType64 = "sugg64"
)

type Suggestion struct {
	Id   int64             `form:"id" json:"id" binding:"required"`
	Name string            `form:"name" json:"name" binding:"required"`
	Type string            `form:"type" json:"type" binding:"required"`
	Data map[string]string `form:"data" json:"data" binding:"required"`
}

var counter int64 = 2
var database = map[int64]Suggestion{
	1: Suggestion{1, "Default 64", SuggestionType64, map[string]string{
		"乾":  "乾，元亨利貞。",
		"坤":  "元亨，利牝馬之貞。君子有攸往，先迷，後得主利。西南得朋，東北喪朋，安貞吉。",
		"屯":  "元亨利貞，勿用有攸往，利建侯。",
		"蒙":  "亨。匪我求童蒙，童蒙求我。初筮告，再三瀆，瀆則不告。利貞。",
		"需":  "有孚，光亨，貞吉，利涉大川。",
		"訟":  "訟，有孚窒惕，中吉，終凶。利見大人，不利涉大川。",
		"師":  "師貞，丈人吉，無咎。",
		"比":  "比，吉，原筮，元永貞，無咎。不寧方來，後夫凶。",
		"小畜": "小畜，亨。密雲不雨，自我西郊。",
		"履":  "履虎尾，不咥人，亨。",
		"泰":  "泰，小往大來，吉亨。",
		"否":  "否之匪人，不利君子貞，大往小來。",
		"同人": "同人於野，亨。利涉大川，利君子貞。",
		"大有": "大有，元亨。",
		"謙":  "謙，亨，君子有終。",
		"豫":  "豫，利建侯、行師。",
		"隨":  "隨，元亨利貞，無咎。",
		"蠱":  "蠱，元亨，利涉大川。先甲三日，後甲三日。",
		"臨":  "臨，元亨利貞，至於八月有凶。",
		"觀":  "觀，盥而不薦，有孚顒若。",
		"噬嗑": "噬嗑，亨，利用獄。",
		"賁":  "賁亨，小利有攸往。",
		"剝":  "剝，不利有攸往。",
		"復":  "復，亨。出入無疾，朋來無咎。反覆其道，七日來復，利有攸往。",
		"無妄": "無妄，元亨利貞。其匪正有眚，不利有攸往。",
		"大畜": "大畜，利貞，不家食吉，利涉大川。",
		"頤":  "頤，貞吉，觀頤，自求口實。",
		"大過": "大過，棟橈。利有攸往，亨。",
		"坎":  "習坎，有孚，維心亨，行有尚。",
		"離":  "離，利貞，亨，畜牝牛，吉。",
		"咸":  "咸亨，利貞，取女吉。",
		"恆":  "恆，亨，無咎，利貞，利有攸往。",
		"遯":  "遯亨，小利貞。",
		"大壯": "大壯，利貞。",
		"晉":  "晉，康侯用錫馬蕃庶，晝日三接。",
		"明夷": "明夷，利艱貞。",
		"家人": "家人，利女貞。",
		"睽":  "睽，小事吉。",
		"蹇":  "蹇，利西南，不利東北，利見大人，貞吉。",
		"解":  "解，利西南，無所往，其來復吉；有攸往，夙吉。",
		"損":  "損，有孚，元吉，無咎可貞，利有攸往。曷之用，二簋可用享。",
		"益":  "益，利有攸往，利涉大川。",
		"夬":  "夬，揚於王庭，孚號有厲。告自邑，不利即戎，利有攸往。",
		"姤":  "姤，女壯，勿用取女。",
		"萃":  "萃亨，王假有廟，利見大人，亨，利貞。用大牲吉。利有攸往。",
		"升":  "升，元亨。用見大人，勿恤。南征吉。",
		"困":  "困，亨，貞大人吉，無咎。有言不信。",
		"井":  "井，改邑不改井，無喪無得，往來井井，汔至，亦未繘井，羸其瓶，凶。",
		"革":  "革，已日乃孚，元亨利貞，悔亡。",
		"鼎":  "鼎，元吉，亨。",
		"震":  "震，亨。震來虩虩，笑言啞啞，震驚百里，不喪匕鬯。",
		"艮":  "艮其背，不獲其身；行其庭，不見其人。",
		"漸":  "漸，女歸吉，利貞。",
		"歸妹": "歸妹，征凶，無攸利。",
		"豐":  "豐亨，王假之，勿憂，宜日中。",
		"旅":  "旅，小亨，旅貞吉。",
		"巽":  "巽，小亨，利有攸往，利見大人。",
		"兌":  "兌，亨，利貞。",
		"渙":  "渙亨，王假有廟，利涉大川，利貞。",
		"節":  "節亨。苦節，不可貞。",
		"中孚": "中孚，豚魚吉，利涉大川，利貞。",
		"小過": "小過，亨，利貞。可小事，不可大事。飛鳥遺之音，不宜上，宜下。大吉。",
		"既濟": "既濟，亨小，利貞。初吉終亂。",
		"未濟": "未濟，亨。小狐汔濟，濡其尾，無攸利。",
	}},
	2: Suggestion{2, "Default 4", SuggestionType4, map[string]string{
		"乾": "乾，元亨利貞。",
		"坤": "元亨，利牝馬之貞。君子有攸往，先迷，後得主利。西南得朋，東北喪朋，安貞吉。",
		"泰": "泰，小往大來，吉亨。",
		"否": "否之匪人，不利君子貞，大往小來。",
	}},
}

func Add(context *gin.Context) {
	type JsonType struct {
		Name string            `form:"name" json:"name" binding:"required"`
		Type string            `form:"type" json:"type" binding:"required"`
		Data map[string]string `form:"data" json:"data" binding:"required"`
	}

	var json JsonType
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if json.Type != SuggestionType4 && json.Type != SuggestionType64 {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if json.Type == SuggestionType4 {
		table := []string{"乾", "坤", "泰", "否"}
		for _, key := range table {
			_, ok := json.Data[key]
			if !ok {
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}
	}

	if json.Type == SuggestionType64 {
		table := utility.GetHexagramSequenceTable()
		for _, key := range table {
			_, ok := json.Data[key]
			if !ok {
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}
	}

	counter += 1
	suggestion := Suggestion{counter, json.Name, json.Type, json.Data}
	database[counter] = suggestion

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      counter,
	})
}

func Get(context *gin.Context) {
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

	suggestionId, ok := utility.QueryInt("suggestion", context)
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	table, ok := database[suggestionId]
	if !ok {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if table.Type != SuggestionType64 && table.Type != SuggestionType4 {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	suggestionArray := []string{}
	if table.Type == SuggestionType64 {
		array, ok := record.GetRecordFrequencyArray(recordId, tableId)
		if !ok {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		sort.Slice(array, func(i, j int) bool {
			return array[i].Value > array[j].Value // 降序
		})
		if table.Type == SuggestionType64 {
			for i, value := range array {
				fmt.Println(i, value)
				suggestionArray = append(suggestionArray, table.Data[value.Hexagram])
				if i >= 2 {
					break
				}
			}
		}
	} else {
		summary := map[string]int{"乾": 0, "坤": 0, "泰": 0, "否": 0}
		freqTable, ok := record.GetFrequencyTable(recordId, tableId)
		if !ok {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		for i, array := range freqTable {
			for j, value := range array {
				if value == 0 {
					continue
				}
				if i < 4 {
					if j < 4 {
						summary["乾"] += value
					} else {
						summary["否"] += value
					}
				} else {
					if j < 4 {
						summary["泰"] += value
					} else {
						summary["坤"] += value
					}
				}
			}
		}

		maxKey := ""
		maxValue := 0
		for key, value := range summary {
			if maxValue > value {
				continue
			}

			maxKey = key
			maxValue = value
		}

		suggestionArray = append(suggestionArray, table.Data[maxKey])

	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    suggestionArray,
	})
}

func GetList(context *gin.Context) {
	type JsonType struct {
		Id   int64  `form:"id" json:"id" binding:"required"`
		Name string `form:"name" json:"name" binding:"required"`
		Type string `form:"type" json:"type" binding:"required"`
	}

	array := []JsonType{}
	for _, value := range database {
		array = append(array, JsonType{value.Id, value.Name, value.Type})
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    array,
	})
}
