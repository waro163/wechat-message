package mp

import (
	"encoding/xml"
	"errors"

	msg "github.com/waro163/wechat-message/mp/message"
)

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Passive_user_reply_message.html

var XmlContentType = []string{"application/xml; charset=utf-8"}
var PlainContentType = []string{"text/plain; charset=utf-8"}

// ErrInvalidReply 无效的回复
var ErrInvalidReply = errors.New("无效的回复消息")

// ErrUnsupportReply 不支持的回复类型
var ErrUnsupportReply = errors.New("不支持的回复消息")

func ReplyMsg(in MixMessage, out interface{}) ([]byte, error) {
	switch reply := out.(type) {
	case *msg.Text:
		reply.SetToUserName(in.FromUserName)
		reply.SetFromUserName(in.ToUserName)
		return xml.Marshal(reply)
	case *msg.Image:
		reply.SetToUserName(in.FromUserName)
		reply.SetFromUserName(in.ToUserName)
		return xml.Marshal(reply)
	case *msg.Voice:
		reply.SetToUserName(in.FromUserName)
		reply.SetFromUserName(in.ToUserName)
		return xml.Marshal(reply)
	case *msg.Video:
		reply.SetToUserName(in.FromUserName)
		reply.SetFromUserName(in.ToUserName)
		return xml.Marshal(reply)
	case *msg.Music:
		reply.SetToUserName(in.FromUserName)
		reply.SetFromUserName(in.ToUserName)
		return xml.Marshal(reply)
	case *msg.News:
		reply.SetToUserName(in.FromUserName)
		reply.SetFromUserName(in.ToUserName)
		return xml.Marshal(reply)
	default:
		return nil, ErrUnsupportReply
	}
}
