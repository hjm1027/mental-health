package message

import (
	"github.com/gin-gonic/gin"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
)

type CountModel struct {
	Count uint32 `json:"count"`
}

//  获取消息提醒的个数
func Count(c *gin.Context) {
	id, _ := c.Get("id")
	count, err := model.GetCount(id.(uint32))
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, CountModel{Count: count})
}
