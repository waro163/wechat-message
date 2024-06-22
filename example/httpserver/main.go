package main

import (
	"net/http"

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
	handle := serve.Server{
		Token:        "",
		SkipValidata: true,
		LogMsg:       true,
		Handler:      &handleMsg{},
	}

	http.HandleFunc("/wechat/callback", handle.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
