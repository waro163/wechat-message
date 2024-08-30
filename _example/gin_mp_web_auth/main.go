package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var appId = ""

func main() {
	r := gin.Default()

	r.Any("/wechat/auth")
	r.Run(":8080")
}

func Auth(c *gin.Context) {
	authCode := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "empty code"})
		return
	}

}
