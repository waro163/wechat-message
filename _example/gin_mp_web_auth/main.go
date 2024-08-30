package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appId     = ""
	appSecret = ""
)

func main() {
	r := gin.Default()

	r.Any("api/wechat/auth", WechatAuth)
	r.Run(":8080")
}

// step1. https://open.weixin.qq.com/connect/oauth2/authorize?appid=[APPID]&redirect_uri=[localhost:8080/api/wechat/auth]&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect
func WechatAuth(c *gin.Context) {
	authCode := c.Query("code")
	if authCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "empty code"})
		return
	}
	tokenUrlFormat := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(tokenUrlFormat, appId, appSecret, authCode)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Printf("get token error, %s", err)
		c.Status(http.StatusOK)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var tokenResp TokenResp
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		log.Printf("json Unmarshal error, %s", &body)
		c.Status(http.StatusOK)
		return
	}
	log.Printf("user open_id: %s", tokenResp.Openid)
	c.Status(http.StatusOK)
}

type TokenResp struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	Openid         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
	Unionid        string `json:"unionid"`
}
