package mini

import (
	"github.com/waro163/wechat-message/mp"
	msg "github.com/waro163/wechat-message/mp/message"
)

type MixMessage struct {
	msg.CommonMsg

	// 基本消息
	// --------text--------
	MsgID   int64  `json:"MsgId" xml:"MsgId"`     //text,
	Content string `json:"Content" xml:"Content"` //text,

	// --------image--------
	// MsgID        int64   `xml:"MsgId"`
	PicURL  string `json:"PicUrl" xml:"PicUrl"`   //image,
	MediaID string `json:"MediaId" xml:"MediaId"` //image,

	// --------mini card--------
	// MsgID        int64   `xml:"MsgId"`
	Title        string `json:"Title" xml:"Title"`
	AppId        string `json:"AppId" xml:"AppId"`
	PagePath     string `json:"PagePath" xml:"PagePath"`
	ThumbUrl     string `json:"ThumbUrl" xml:"ThumbUrl"`
	ThumbMediaId string `json:"ThumbMediaId" xml:"ThumbMediaId"`

	// 事件相关
	// --------进入客服会话--------
	Event       EventType `json:"Event" xml:"Event"`
	SessionFrom string    `json:"SessionFrom" xml:"SessionFrom"`

	// --------订阅消息事件推送--------
	// Event EventType `json:"Event" xml:"Event"`
	SubscribeMsgPopupEventList *[]mp.SubscribeMsgPopupEvent `json:"List" xml:"SubscribeMsgPopupEvent>List"`

	// --------管理订阅--------
	// Event EventType `xml:"Event"`
	SubscribeMsgChangeEvent *struct {
		TemplateID            string `xml:"TemplateId"`
		SubscribeStatusString string `xml:"SubscribeStatusString"`
	} `json:"List" xml:"SubscribeMsgChangeEvent>List"`

	// --------订阅消息发送结果--------
	// Event EventType `xml:"Event"`
	SubscribeMsgSentEvent *struct {
		TemplateID  string `xml:"TemplateId"`
		MsgID       int64  `xml:"MsgID"`
		ErrorCode   int32  `xml:"ErrorCode"`
		ErrorStatus string `xml:"ErrorStatus"`
	} `json:"List" xml:"SubscribeMsgSentEvent>List"` // subscribe_msg_sent_event
}
