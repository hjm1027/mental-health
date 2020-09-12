package reserve

import (
	"strconv"
	"time"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

/*
type ReserveQueryRequest struct {
	Weekday  uint8 `json:"weekday" binding:"required"`
	Schedule uint8 `json:"schedule" binding:"required"`
}*/

type ReserveQueryResponse struct {
	CanReserve bool `json:"can_reserve"`
}

//查询所选时间是否可以预约
func QueryReserve(c *gin.Context) {
	weekStr := c.Query("weekday")
	weekday, err := strconv.ParseUint(weekStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	schStr := c.Query("schedule")
	schedule, err := strconv.ParseUint(schStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	time := time.Now().UTC().Add(8 * time.Hour)

	canReserve, err := model.QueryReserve(uint8(weekday), uint8(schedule), time)
	if err != nil {
		handler.SendError(c, errno.ErrQueryReserve, nil, err.Error())
		return
	}

	response := ReserveQueryResponse{
		CanReserve: canReserve,
	}

	handler.SendResponse(c, errno.OK, response)
}
