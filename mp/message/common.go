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
