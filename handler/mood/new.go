package mood

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type MoodRequest struct {
	Score uint8  `json:"score" binding:"required"`
	Note  string `json:"note" binding:"required"`
}

func mix(year uint32, month uint8, day uint8) string {
	var monthstr, daystr string
	if month < 10 {
		monthstr = "0" + strconv.Itoa(int(month))
	} else {
		monthstr = strconv.Itoa(int(month))
	}
	if day < 10 {
		daystr = "0" + strconv.Itoa(int(day))
	} else {
		daystr = strconv.Itoa(int(day))
	}
	str := strconv.Itoa(int(year)) + "." + monthstr + "." + daystr
	return str
}

//添加心情信息
func NewMood(c *gin.Context) {
	var data MoodRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	//type Month int
	//const (January Month = 1 + iota)
	year := uint32(time.Now().UTC().Add(8 * time.Hour).Year())
	month := uint8(time.Now().UTC().Add(8 * time.Hour).Month())
	day := uint8(time.Now().UTC().Add(8 * time.Hour).Day())

	//monthmap := map[string]uint8{"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6,
	//	"July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12}
	//month := monthmap[monthstr]

	var mood = &model.MoodModel{
		UserId: userId,
		Date:   mix(year, month, day),
		Year:   year,
		Month:  month,
		Day:    day,
		Score:  data.Score,
		Note:   data.Note,
	}
	fmt.Println(mood.Date)

	if err := mood.Have(); err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}

	if err := mood.New(); err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}

	handler.SendResponse(c, errno.OK, nil)
}
