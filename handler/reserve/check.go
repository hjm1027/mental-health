package reserve

import (
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type ReserveCheckRequest struct {
	Weekday  uint8 `json:"weekday" binding:"required"`
	Schedule uint8 `json:"schedule" binding:"required"`
	Check    bool  `json:"check"`
}

//审核预约
func CheckReserve(c *gin.Context) {
	userId := c.MustGet("id").(uint32)
	var data ReserveCheckRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	/*time2 := time.Now().UTC().Add(8 * time.Hour)
	teacher, err := model.GetTeacherBySchedule(data.Weekday, data.Schedule)
	if err != nil {
		handler.SendError(c, errno.ErrGetTeacherBySchedule, nil, err.Error())
		return
	}*/

	//获取预约记录
	record, err := model.GetReserveRecord(userId, data.Weekday, data.Schedule)
	if err != nil {
		handler.SendError(c, errno.GetReserveRecord, nil, err.Error())
		return
	}

	var status uint8
	if data.Check {
		status = 2
		if err := record.UpdateRecord(status); err != nil {
			handler.SendError(c, errno.ErrUpdateRecord, nil, err.Error())
			return
		}
	} else {
		status = 0
		if err := record.UpdateRecord(status); err != nil {
			handler.SendError(c, errno.ErrUpdateRecord, nil, err.Error())
			return
		}
	}

	if err := model.CheckReserve(data.Weekday, data.Schedule, status); err != nil {
		handler.SendError(c, errno.ErrCheckReserve, nil, err.Error())
		return
	}

	handler.SendResponse(c, errno.OK, nil)
}
