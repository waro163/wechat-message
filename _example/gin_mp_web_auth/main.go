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
	appId     = "your_app_id"
	appSecret = "your_app_secret"
)

func main() {
	r := gin.Default()
	r.Any("/MP_verify_xxxx.txt", BindWebAuthHost)
	r.Any("/api/wechat/callback_bind", BindCallback)
	r.Any("/api/wechat/auth", WechatAuth)
	r.Run(":8080")
}

func BindWebAuthHost(c *gin.Context) {
	c.String(http.StatusOK, "aaaa")
}

func BindCallback(c *gin.Context) {
	echostr := c.Query("echostr")
	log.Printf("echostr: %s\n", echostr)
	c.String(http.StatusOK, echostr)
}

// step0. 注意，测试号管理中，在接口权限表 - 网页账号 处修改该域名，并非JS接口安全域名处修改，切记
// step1. https://open.weixin.qq.com/connect/oauth2/authorize?appid=[your_app_id]&redirect_uri=[your_domain]/api/wechat/auth&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect
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
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var tokenResp TokenResp
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		log.Printf("json Unmarshal error, %s", &body)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	log.Printf("user open_id: %s", tokenResp.Openid)
	c.JSON(http.StatusOK, gin.H{"open_id": tokenResp.Openid})
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
