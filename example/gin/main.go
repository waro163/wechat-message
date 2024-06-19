package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waro163/wechat-message/mp"
	"github.com/waro163/wechat-message/mp/message"
	"github.com/waro163/wechat-message/mp/serve"
)

func main() {
	r := gin.Default()

	handle := serve.Server{
		Token:        "",
		SkipValidata: true,
		LogMsg:       true,
		Handler: func(mm mp.MixMessage) (interface{}, error) {
			text := message.NewText("you got it!")
			return text, nil
		},
	}
	r.Any("/wechat/callback", gin.WrapH(&handle))
	r.Run(":8080")
}
