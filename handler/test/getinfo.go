package test

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	test := model.TestModel{Id: uint32(id)}
	if err := test.GetById(); err != nil {
		handler.SendError(c, errno.ErrGetTestInfo, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, test)
}
