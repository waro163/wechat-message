package message

import "encoding/xml"

// MsgType 基本消息类型
type MsgType string

// MarshalXML 实现序列化方法
func (mt MsgType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(mt)}, start)
}

const (
	// MsgTypeText 表示文本消息
	MsgTypeText MsgType = "text"
	// MsgTypeImage 表示图片消息
	MsgTypeImage MsgType = "image"
	// MsgTypeVoice 表示语音消息
	MsgTypeVoice MsgType = "voice"
	// MsgTypeVideo 表示视频消息
	MsgTypeVideo MsgType = "video"
	// MsgTypeMiniprogrampage 表示小程序卡片消息
	MsgTypeMiniprogrampage MsgType = "miniprogrampage"
	// MsgTypeShortVideo 表示短视频消息[限接收]
	MsgTypeShortVideo MsgType = "shortvideo"
	// MsgTypeLocation 表示坐标消息[限接收]
	MsgTypeLocation MsgType = "location"
	// MsgTypeLink 表示链接消息[限接收]
	MsgTypeLink MsgType = "link"
	// MsgTypeMusic 表示音乐消息[限回复]
	MsgTypeMusic MsgType = "music"
	// MsgTypeNews 表示图文消息[限回复]
	MsgTypeNews MsgType = "news"
	// MsgTypeTransfer 表示消息消息转发到客服[限回复]
	MsgTypeTransfer MsgType = "transfer_customer_service"
	// MsgTypeEvent 表示事件推送消息
	MsgTypeEvent MsgType = "event"
)
