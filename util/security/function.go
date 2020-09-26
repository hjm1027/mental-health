package security

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type accessTokenManager struct {
	Token     string
	CreateAt  time.Time
	ExpiresIn time.Duration
}

var (
	AppID     string
	AppSecret string

	accessToken = &accessTokenManager{ExpiresIn: 7200 * time.Second}

	msgSecCheckURL    = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token="
	accessTokenGetURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v"
)

func WXSecInit() {
	AppID = viper.GetString("app.app_id")
	AppSecret = viper.GetString("app.app_secret")

	err := accessToken.loadToken()
	if err != nil {
		log.Error("accessToken.loadToken() error", err)
		return
	}

	msgSecCheckURL += accessToken.Token
}

type WXGetTokenPayload struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
	ErrCode     int32  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func (t *accessTokenManager) loadToken() error {
	fmt.Println(fmt.Sprintf(accessTokenGetURL, AppID, AppSecret))
	resp, err := http.Get(fmt.Sprintf(accessTokenGetURL, AppID, AppSecret))
	//fmt.Println(resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var obj WXGetTokenPayload
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		log.Error("json unmarshal to WXGetTokenPayload error", err)
		return err
	}

	//fmt.Printf("WX access token: old token: %s; new token: %s\n", t.Token, obj.AccessToken)

	t.Token = obj.AccessToken
	t.CreateAt = time.Now().UTC().Add(8 * time.Hour)
	//fmt.Println(t)
	//fmt.Println(accessToken)

	return nil
}

/*
func (t *accessTokenManager) check() error {
	now := time.Now().UTC().Add(8 * time.Hour)
	if t.CreateAt.Add(t.ExpiresIn).Sub(now) <= 0 {
		// 过期，更新 token
		if err := t.loadToken(); err != nil {
			log.Error("Refresh access token failed", err)
			return err
		}
		log.Info("Refresh access token OK")
	}

	//fmt.Printf("WX access token info: createAt=%v, expiresIn=%v, sub time from now=%v\n",
	//	t.CreateAt, t.ExpiresIn, t.CreateAt.Add(t.ExpiresIn).Sub(now))

	return nil
}*/

func RefreshTokenScheduled() {
	for {
		// 提前60分钟更新
		time.Sleep(accessToken.ExpiresIn - time.Minute*60)
		//time.Sleep(accessToken.ExpiresIn - time.Second*7190)

		if err := accessToken.loadToken(); err != nil {
			log.Error("Refresh access token failed", err)
			continue
		}

		log.Info("Refresh WX access token OK")
	}
}
