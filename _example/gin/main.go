package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waro163/wechat-message/mp"
	"github.com/waro163/wechat-message/mp/message"
	"github.com/waro163/wechat-message/mp/serve"
)

type handleMsg struct{}

func (h *handleMsg) HandleMsg(mp.MixMessage) (interface{}, error) {
	text := message.NewText("you got it!")
	return text, nil
}

func main() {
	r := gin.Default()

	handle := serve.Server{
		Token:        "",
		SkipValidata: true,
		LogMsg:       true,
		Handler:      &handleMsg{},
	}
	r.Any("/wechat/callback", gin.WrapH(&handle))
	r.Run(":8080")
}
