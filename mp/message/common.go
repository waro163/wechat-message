package message

import "encoding/xml"

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

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
	// MsgTypeTransfer 表示消息消息转发到客服
	MsgTypeTransfer MsgType = "transfer_customer_service"
	// MsgTypeEvent 表示事件推送消息
	MsgTypeEvent MsgType = "event"
)

// CommonMsg 消息中通用的结构
type CommonMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName" json:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime" json:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType" json:"MsgType"`
}

// SetToUserName set ToUserName
func (msg *CommonMsg) SetToUserName(toUserName CDATA) {
	msg.ToUserName = toUserName
}

// SetFromUserName set FromUserName
func (msg *CommonMsg) SetFromUserName(fromUserName CDATA) {
	msg.FromUserName = fromUserName
}

// SetCreateTime set createTime
func (msg *CommonMsg) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

// SetMsgType set MsgType
func (msg *CommonMsg) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}

// GetOpenID get the FromUserName value
func (msg *CommonMsg) GetOpenID() string {
	return string(msg.FromUserName)
}
