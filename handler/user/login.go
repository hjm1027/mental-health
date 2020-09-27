package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/pkg/token"
	"github.com/mental-health/util"
	"github.com/spf13/viper"
)

type CodeResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrCode    int32  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

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
			IsTeacher:    false,
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

	//获取用户openid
	AppID := viper.GetString("app.app_id")
	AppSecret := viper.GetString("app.app_secret")

	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	code2Session := fmt.Sprintf(url, AppID, AppSecret, l.Jscode)

	resp, err := http.Get(code2Session)
	if err != nil {
		handler.SendError(c, errno.ErrGetUserOpenid, nil, err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handler.SendError(c, errno.ErrGetUserOpenid, nil, err.Error())
		return
	}

	var rp CodeResponse
	if err := json.Unmarshal(body, &rp); err != nil {
		handler.SendError(c, errno.ErrGetUserOpenid, nil, err.Error())
		return
	}

	if rp.ErrCode != 0 {
		log.Info(fmt.Sprintf("get user openid failed. code: %d; msg: %s.", rp.ErrCode, rp.ErrMsg))
		handler.SendError(c, errno.ErrGetUserOpenid, nil, rp.ErrMsg)
		return
	}
	log.Info("get user openid OK.")

	//保存openid
	have, err := model.HaveUserCode(u.Id)
	if err != nil {
		handler.SendError(c, errno.ErrSaveUserOpenid, nil, err.Error())
		return
	}

	code := &model.UserCodeModel{
		UserId:  u.Id,
		Openid:  rp.Openid,
		Unionid: rp.Unionid,
	}

	if have == 0 {
		if err := code.Create(); err != nil {
			handler.SendError(c, errno.ErrSaveUserOpenid, nil, err.Error())
			return
		}
	} else {
		if err := code.Save(); err != nil {
			handler.SendError(c, errno.ErrSaveUserOpenid, nil, err.Error())
			return
		}
	}
	log.Info("save user openid OK.")

	handler.SendResponse(c, errno.OK, model.AuthResponse{Token: t, IsNew: IsNewUser})
}
