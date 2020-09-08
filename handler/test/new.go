package test

import (
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type NewTestRequest struct {
	Url     string `json:"url" binding:"required"`
	Header  string `json:"header" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type NewTestResponse struct {
	TestId uint32 `json:"test_id"`
}

func New(c *gin.Context) {
	var data NewTestRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	var test = &model.TestModel{
		Url:     data.Url,
		Header:  data.Header,
		Content: data.Content,
	}

	if err := test.New(); err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}
	response := NewTestResponse{TestId: test.Id}

	handler.SendResponse(c, nil, response)
}
