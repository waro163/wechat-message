package serve

import "github.com/waro163/wechat-message/mp"

type IMessageHandler interface {
	HandleMsg(mp.MixMessage) (interface{}, error)
}

type customMessagehandler struct {
	msgHandler func(mp.MixMessage) (interface{}, error)
}

func NewCustomMessageHandler(msgHandler func(mp.MixMessage) (interface{}, error)) IMessageHandler {
	return &customMessagehandler{
		msgHandler: msgHandler,
	}
}

func (h *customMessagehandler) HandleMsg(msg mp.MixMessage) (interface{}, error) {
	return h.msgHandler(msg)
}
