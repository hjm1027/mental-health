package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
)

func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var l model.LoginModel
	if err := c.ShouldBindJSON(&l); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	handler.SendResponse(c, errno.OK, nil)
}
