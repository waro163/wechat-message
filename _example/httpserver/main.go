package main

import (
	"net/http"

	"github.com/waro163/wechat-message/mp"
	"github.com/waro163/wechat-message/mp/message"
	"github.com/waro163/wechat-message/mp/serve"
)

var msgHander = serve.NewCustomMessageHandler(func(mm mp.MixMessage) (interface{}, error) {
	text := message.NewText("you got it!")
	return text, nil
})

func main() {
	handle := serve.Server{
		Token:        "",
		SkipValidata: true,
		LogMsg:       true,
		Handler:      msgHander,
	}

	http.HandleFunc("/wechat/callback", handle.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
