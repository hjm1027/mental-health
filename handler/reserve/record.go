package reserve

import (
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

//获取我的记录
func GetRecord(c *gin.Context) {
	userId := c.MustGet("id").(uint32)

	records, err := model.GetRecords(userId)
	if err != nil {
		handler.SendError(c, errno.ErrGetRecord, nil, err.Error())
		return
	}

	handler.SendResponse(c, errno.OK, records)
}
