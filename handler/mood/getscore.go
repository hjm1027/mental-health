package mood

import (
	"strconv"
	"time"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type MoodScoreRequest struct {
	Year uint32 `json:"year" binding:"required"`
}

type MoodScoreResponse struct {
	Month uint8                 `json:"month"`
	List  []model.MoodScoreItem `json:"list"`
}

//获取一年的心情指数
func GetMoodScore(c *gin.Context) {
	userId := c.MustGet("id").(uint32)
	yearstr := c.Query("year")
	year, _ := strconv.ParseInt(yearstr, 10, 64)
	month := uint8(time.Now().UTC().Add(8 * time.Hour).Month())
	/*monthstr := c.Query("month")
	month, _ := strconv.ParseInt(monthstr, 10, 64)*/

	var i uint8
	if year == int64(time.Now().UTC().Add(8*time.Hour).Year()) {
		i = uint8(month)
	} else {
		i = 12
	}

	//var finResponse MoodScoreResponseList
	var finResponse []MoodScoreResponse
	for ; i > 0; i-- {
		response, err := model.GetMoodScore(userId, uint32(year), uint8(i))
		if err != nil {
			handler.SendError(c, errno.ErrGetScoreInfo, nil, err.Error())
		}
		scorelist := MoodScoreResponse{
			Month: uint8(i),
			List:  response,
		}
		finResponse = append(finResponse, scorelist)
	}

	handler.SendResponse(c, errno.OK, finResponse)
}
