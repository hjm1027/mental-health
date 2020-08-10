package message

import (
	"github.com/MuxiKeStack/muxiK-StackBackend/model"

	"github.com/gin-gonic/gin"
)

type messageListResponse struct {
	messageList *[]model.MessageSub
}

func Get(c *gin.Context) { /*
		userId := c.MustGet("id").(uint32)

		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.ParseUint(pageStr, 10, 32)
		if err != nil {
			handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
			return
		}

		limitStr := c.DefaultQuery("limit", "10")
		limit, err := strconv.ParseUint(limitStr, 10, 32)
		if err != nil {
			handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
			return
		}

		messageList, err := service.MessageList(uint32(page), uint32(limit), userId)
		if err != nil {
			handler.SendError(c, errno.ErrGetMessage, nil, err.Error())
			return
		}
		handler.SendResponse(c, nil, messageList)*/
}
