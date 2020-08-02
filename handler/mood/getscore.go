package mood

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type MoodScoreRequest struct {
	Year  uint32 `json:"year" binding:"required"`
	Month uint8  `json:"month" binding:"required"`
}

type MoodScoreResponse struct {
	Month uint8                 `json:"month"`
	List  []model.MoodScoreItem `json:"list"`
}

type MoodScoreResponseList struct {
	List []MoodScoreResponse `json:"list"`
}

//获取一个月的心情指数
func GetMoodScore(c *gin.Context) { /*
		var data MoodScoreRequest
		if err := c.ShouldBindJSON(&data); err != nil {
			handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
			return
		}*/

	userId := c.MustGet("id").(uint32)
	yearstr := c.Query("year")
	year, _ := strconv.ParseInt(yearstr, 10, 64)
	monthstr := c.Query("month")
	month, _ := strconv.ParseInt(monthstr, 10, 64)

	//var finResponse MoodScoreResponseList
	var finResponse []MoodScoreResponse
	for i := month; i > 0; i-- {
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
