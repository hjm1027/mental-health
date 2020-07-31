package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/pkg/token"
	"github.com/mental-health/util"
	"github.com/spf13/viper"
)

func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var l model.LoginModel
	if err := c.ShouldBindJSON(&l); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// Compare the login password with the user password.
	// 业务逻辑异常，使用 SendResponse 发送 200 请求 + 自定义错误码
	if err := util.LoginRequest(l.Sid, l.Password); err != nil {
		handler.SendResponse(c, errno.ErrAuthFailed, nil)
		return
	}

	// judge whether there is this user or not
	user := model.UserModel{Sid: l.Sid}
	IsNewUser, _ := user.HaveUser()
	if IsNewUser == 0 {
		err := user.Create()
		if err != nil {
			handler.SendError(c, errno.ErrCreateUser, nil, err.Error())
			return
		}
	}

	// get user
	u, err := model.GetUserBySid(l.Sid)
	if err != nil {
		handler.SendError(c, errno.ErrUserNotFound, nil, err.Error())
		return
	}

	if IsNewUser == 0 {
		err := u.UpdateInfo(&model.UserInfoRequest{
			Avatar:       viper.GetString("default_user.avatar"),
			Username:     viper.GetString("default_user.username") + strconv.FormatUint(uint64(u.Id), 10),
			Introduction: viper.GetString("default_user.introduction"),
			Phone:        viper.GetString("default_user.phone"),
			Back_avatar:  viper.GetString("default_user.back_avatar"),
		})
		if err != nil {
			handler.SendError(c, errno.ErrUpdateUser, nil, err.Error())
			return
		}
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{Id: u.Id}, "")
	if err != nil {
		handler.SendError(c, errno.ErrToken, nil, err.Error())
		return
	}

	handler.SendResponse(c, errno.OK, model.AuthResponse{Token: t, IsNew: IsNewUser})
}
