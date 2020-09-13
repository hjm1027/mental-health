package reserve

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

//获取记录的详细信息
func GetRecordInfo(c *gin.Context) {
	//userId := c.MustGet("id").(uint32)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	record := model.RecordModel{Id: uint32(id)}
	if err := record.GetInfo(); err != nil {
		handler.SendError(c, errno.ErrGetRecordInfo, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, record)
}
