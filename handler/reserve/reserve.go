package reserve

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
)

type ReserveRequest struct {
	Weekday  uint8  `json:"weekday" binding:"required"`
	Schedule uint8  `json:"schedule" binding:"required"`
	Type     uint8  `json:"type"`
	Method   uint8  `json:"method"`
	ModuleId uint32 `json:"module_id"`
}

func getAdvanceTime(weekday, weekdayNow uint8) uint8 {
	difference := map[int]int{-6: 8, -5: 2, -4: 3, -3: 4, -2: 5, -1: 6, 0: 7, 1: 8, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	advance := difference[int(weekday)-int(weekdayNow)]
	return uint8(advance)
}

//进行预约
func Reserve(c *gin.Context) {
	userId := c.MustGet("id").(uint32)

	var data ReserveRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	time2 := time.Now().UTC().Add(8 * time.Hour)
	weekday := time.Now().UTC().Add(8 * time.Hour).Weekday()

	teacher, err := model.GetTeacherBySchedule(data.Weekday, data.Schedule)
	if err != nil {
		handler.SendError(c, errno.ErrGetTeacherBySchedule, nil, err.Error())
		return
	}

	advanceTime := getAdvanceTime(data.Weekday, uint8(weekday))

	reserve := &model.ReserveModel{
		Weekday:     data.Weekday,
		Schedule:    data.Schedule,
		Teacher:     teacher,
		Reserve:     1,
		Time:        time2,
		AdvanceTime: advanceTime,
		Type:        data.Type,
		Method:      data.Method,
		UserId:      userId,
	}

	err1, err2 := reserve.New(userId)
	if err1 != nil {
		handler.SendError(c, errno.ErrCreateReserve, nil, err1.Error())
		return
	}
	if err2 != nil {
		handler.SendError(c, errno.ErrCreateReserve, nil, err2.Error())
		return
	}

	handler.SendResponse(c, errno.OK, nil)
}
