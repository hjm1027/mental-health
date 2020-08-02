package mood

import (
	"strconv"
	"time"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type MoodNoteRequest struct {
	Year  uint32 `json:"year" binding:"required"`
	Month uint8  `json:"month" binding:"required"`
}

type MoodNoteResponse struct {
	Month uint8                `json:"month"`
	List  []model.MoodNoteItem `json:"list"`
}

//获取一年的心情笔记
func GetMoodNote(c *gin.Context) {
	userId := c.MustGet("id").(uint32)
	yearstr := c.Query("year")
	year, _ := strconv.ParseInt(yearstr, 10, 64)
	month := uint8(time.Now().Month())

	var i uint8
	if year == int64(time.Now().Year()) {
		i = month
	} else {
		i = 12
	}

	//var finResponse MoodNoteResponseList
	var finResponse []MoodNoteResponse
	for ; i > 0; i-- {
		response, err := model.GetMoodNote(userId, uint32(year), uint8(i))
		if err != nil {
			handler.SendError(c, errno.ErrGetNoteInfo, nil, err.Error())
		}
		notelist := MoodNoteResponse{
			Month: uint8(i),
			List:  response,
		}
		finResponse = append(finResponse, notelist)
	}

	handler.SendResponse(c, errno.OK, finResponse)
}
