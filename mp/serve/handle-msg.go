package serve

import "github.com/waro163/wechat-message/mp"

type IMessageHandler interface {
	HandleMsg(mp.MixMessage) (interface{}, error)
}
