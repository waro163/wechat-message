package mp

import (
	msg "github.com/waro163/wechat-message/mp/message"
)

// EventPic 发图事件推送
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

// SubscribeMsgPopupEvent 订阅通知事件推送的消息体
type SubscribeMsgPopupEvent struct {
	TemplateID            string `xml:"TemplateId" json:"TemplateId"`
	SubscribeStatusString string `xml:"SubscribeStatusString" json:"SubscribeStatusString"`
	PopupScene            int    `xml:"PopupScene" json:"PopupScene,string"`
}

// MixMessage 存放所有微信公众号发送过来的消息和事件
type MixMessage struct {
	msg.CommonMsg

	// 基本消息
	MsgID        int64   `xml:"MsgId"`        //text,image,voice,video,location,link
	Content      string  `xml:"Content"`      //text,
	Recognition  string  `xml:"Recognition"`  //voice,
	PicURL       string  `xml:"PicUrl"`       //image,
	MediaID      string  `xml:"MediaId"`      //image,voice,video
	Format       string  `xml:"Format"`       //voice,
	ThumbMediaID string  `xml:"ThumbMediaId"` //video,
	LocationX    float64 `xml:"Location_X"`   //location,
	LocationY    float64 `xml:"Location_Y"`   //location,
	Scale        float64 `xml:"Scale"`        //location,
	Label        string  `xml:"Label"`        //location,
	Title        string  `xml:"Title"`        //link
	Description  string  `xml:"Description"`  //link
	URL          string  `xml:"Url"`          //link

	// 事件相关
	TemplateMsgID int64     `xml:"MsgID"` // 模板消息推送成功的消息是MsgID, template
	Event         EventType `xml:"Event" json:"Event"`
	EventKey      string    `xml:"EventKey"`  //scan_qr,click_menu,
	Ticket        string    `xml:"Ticket"`    //scan,subscribe
	Latitude      string    `xml:"Latitude"`  //report_location
	Longitude     string    `xml:"Longitude"` //report_location
	Precision     string    `xml:"Precision"` //report_location
	MenuID        string    `xml:"MenuId"`    //view-mini
	Status        string    `xml:"Status"`    //template

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"` //scan_push

	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	} `xml:"SendLocationInfo"`

	SubscribeMsgPopupEventList []SubscribeMsgPopupEvent `xml:"SubscribeMsgPopupEvent>List"` //subscribe_msg_popup_event

	SubscribeMsgChangeEvent struct {
		TemplateID            string `xml:"TemplateId"`
		SubscribeStatusString string `xml:"SubscribeStatusString"`
	} `xml:"SubscribeMsgChangeEvent>List"` // subscribe_msg_change_event

	SubscribeMsgSentEvent struct {
		TemplateID  string `xml:"TemplateId"`
		MsgID       int64  `xml:"MsgID"`
		ErrorCode   int32  `xml:"ErrorCode"`
		ErrorStatus string `xml:"ErrorStatus"`
	} `xml:"SubscribeMsgSentEvent>List"` // subscribe_msg_sent_event
}
