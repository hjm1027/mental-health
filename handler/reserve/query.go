package reserve

import (
	"time"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type ReserveQueryRequest struct {
	Weekday  uint8 `json:"weekday" binding:"required"`
	Schedule uint8 `json:"schedule" binding:"required"`
}

type ReserveQueryResponse struct {
	CanReserve bool `json:"can_reserve"`
}

//查询所选时间是否可以预约
func QueryReserve(c *gin.Context) {
	var data ReserveQueryRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	time := time.Now().UTC().Add(8 * time.Hour)

	canReserve, err := model.QueryReserve(data.Weekday, data.Schedule, time)
	if err != nil {
		handler.SendError(c, errno.ErrQueryReserve, nil, err.Error())
		return
	}

	response := ReserveQueryResponse{
		CanReserve: canReserve,
	}

	handler.SendResponse(c, errno.OK, response)
}
