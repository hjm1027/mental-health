package user

import (
	. "github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

//获取用户信息
func GetInfo(c *gin.Context) {
	// log.Info("GetInfo function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := c.Get("id")
	info, err := model.GetUserInfoById(id.(uint32))
	if err != nil {
		SendError(c, errno.ErrGetUserInfo, nil, err.Error())
		return
	}
	SendResponse(c, errno.OK, model.UserInfoResponse{
		Username:     info.Username,
		Avatar:       info.Avatar,
		Sid:          info.Sid,
		Introduction: info.Introduction,
		Phone:        info.Phone,
		Back_avatar:  info.Back_avatar,
		IsTeacher:    info.IsTeacher,
	})
}
