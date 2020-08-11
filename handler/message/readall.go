package message

import (
	"github.com/gin-gonic/gin"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
)

//读取消息提醒
func ReadAll(c *gin.Context) {
	id, _ := c.Get("id")
	err := model.ReadAll(id.(uint32))
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}
	handler.SendResponse(c, errno.OK, nil)
}
