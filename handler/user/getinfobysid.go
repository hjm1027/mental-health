package user

import (
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

//通过学号获取用户信息
func GetInfoBySid(c *gin.Context) {
	//userId := c.MustGet("id").(uint32)
	sid := c.Param("id")
	user := &model.UserModel{Sid: sid}
	if err := user.GetUserBySid(); err != nil {
		handler.SendError(c, errno.ErrGetUserInfo, nil, err.Error())
		return
	}

	response := model.UserInfoResponse{
		Username:     user.Username,
		Avatar:       user.Avatar,
		Sid:          user.Sid,
		Introduction: user.Introduction,
		Phone:        user.Phone,
		Back_avatar:  user.Back_avatar,
		IsTeacher:    user.IsTeacher,
	}

	handler.SendResponse(c, errno.OK, response)
}
