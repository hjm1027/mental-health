package user

import (
	. "github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

func PostInfo(c *gin.Context) {
	// log.Info("PostInfo function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var info model.UserInfoRequest
	// BindJSON 如果字段不存在会返回400
	// ShouldBindJSON 不会自动返回400
	if err := c.BindJSON(&info); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if info.Username == "" || info.Avatar == "" || info.Introduction == "" || info.Phone == "" || info.Back_avatar == "" {
		SendBadRequest(c, errno.ErrUserInfo, nil, "username,avatar,introduction,phone and back_avatar cannot be empty")
		return
	}

	id, _ := c.Get("id")
	if err := model.UpdateInfoById(id.(uint32), &info); err != nil {
		SendError(c, errno.ErrUpdateUser, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
