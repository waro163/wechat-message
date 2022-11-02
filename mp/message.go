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

// MassJobResult 群发消息事件推送
type MassJobResult struct {
	ArticleIdx            int    `xml:"ArticleIdx"`
	UserDeclareState      int32  `xml:"UserDeclareState"`
	AuditState            int32  `xml:"AuditState"`
	OriginalArticleUrl    string `xml:"OriginalArticleUrl"`
	OriginalArticleType   int32  `xml:"OriginalArticleType"`
	CanReprint            int32  `xml:"CanReprint"`
	NeedReplaceContent    int32  `xml:"NeedReplaceContent"`
	NeedShowReprintSource int32  `xml:"NeedShowReprintSource"`
}

// MixMessage 存放所有微信公众号发送过来的消息和事件
type MixMessage struct {
	msg.CommonMsg

	// 基本消息
	// --------text--------
	MsgID   int64  `xml:"MsgId"`   //text,
	Content string `xml:"Content"` //text,

	// --------image--------
	// MsgID        int64   `xml:"MsgId"`
	PicURL  string `xml:"PicUrl"`  //image,
	MediaID string `xml:"MediaId"` //image

	// --------voice--------
	// MsgID        int64   `xml:"MsgId"`
	// MediaID      string  `xml:"MediaId"`//voice
	Format      string `xml:"Format"`      //voice,
	Recognition string `xml:"Recognition"` //voice,

	// --------video/shortvideo--------
	// MsgID        int64   `xml:"MsgId"`
	// MediaID      string  `xml:"MediaId"`  //video
	ThumbMediaID string `xml:"ThumbMediaId"` //video,

	// --------location--------
	// MsgID        int64   `xml:"MsgId"`
	LocationX float64 `xml:"Location_X"` //location,
	LocationY float64 `xml:"Location_Y"` //location,
	Scale     float64 `xml:"Scale"`      //location,
	Label     string  `xml:"Label"`      //location,

	// --------link--------
	// MsgID        int64   `xml:"MsgId"`
	Title       string `xml:"Title"`       //link
	Description string `xml:"Description"` //link
	URL         string `xml:"Url"`         //link

	// 事件相关
	// --------关注/取消关注--------
	Event EventType `xml:"Event" json:"Event"`

	// --------scan qr--------
	// Event         EventType `xml:"Event"`
	EventKey string `xml:"EventKey"` //scan_qr,click_menu,
	Ticket   string `xml:"Ticket"`   //scan,subscribe

	// --------report_location--------
	// Event         EventType `xml:"Event"`
	Latitude  string `xml:"Latitude"`  //report_location
	Longitude string `xml:"Longitude"` //report_location
	Precision string `xml:"Precision"` //report_location

	// --------custom menu--------
	// Event EventType `xml:"Event"`
	// EventKey string `xml:"EventKey"`

	// --------template message--------
	// Event EventType `xml:"Event"`
	TemplateMsgID int64  `xml:"MsgID"`  //MsgID
	Status        string `xml:"Status"` //template

	// --------menu view link/mini--------
	// Event EventType `xml:"Event"`
	// EventKey string `xml:"EventKey"`
	MenuID string `xml:"MenuId"` //view-mini

	// --------scan code push--------
	// Event EventType `xml:"Event"`
	// EventKey string `xml:"EventKey"
	ScanCodeInfo *struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"` //scan_push

	// --------pic--------
	// Event EventType `xml:"Event"`
	// EventKey string `xml:"EventKey"
	SendPicsInfo *struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	// --------location select--------
	// Event EventType `xml:"Event"`
	// EventKey string `xml:"EventKey"
	SendLocationInfo *struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	} `xml:"SendLocationInfo"`

	// --------订阅弹窗--------
	// Event EventType `xml:"Event"`
	SubscribeMsgPopupEventList *[]SubscribeMsgPopupEvent `xml:"SubscribeMsgPopupEvent>List"` //subscribe_msg_popup_event

	// --------管理订阅--------
	// Event EventType `xml:"Event"`
	SubscribeMsgChangeEvent *struct {
		TemplateID            string `xml:"TemplateId"`
		SubscribeStatusString string `xml:"SubscribeStatusString"`
	} `xml:"SubscribeMsgChangeEvent>List"` // subscribe_msg_change_event

	// --------bizsend接口发送--------
	// Event EventType `xml:"Event"`
	SubscribeMsgSentEvent *struct {
		TemplateID  string `xml:"TemplateId"`
		MsgID       int64  `xml:"MsgID"`
		ErrorCode   int32  `xml:"ErrorCode"`
		ErrorStatus string `xml:"ErrorStatus"`
	} `xml:"SubscribeMsgSentEvent>List"` // subscribe_msg_sent_event

	// --------发布能力--------
	PublishEventInfo *struct {
		PublishID     int64  `xml:"publish_id"`
		PublishStatus uint   `xml:"publish_status"`
		ArticleID     string `xml:"article_id"`
		ArticleDetail struct {
			Count uint `xml:"count"`
			Item  []struct {
				Index      uint   `xml:"idx"`
				ArticleURL string `xml:"article_url"`
			} `xml:"item"`
		} `xml:"article_detail"`
		FailIndex []uint `xml:"fail_idx"` //审核不通过字段
	} `xml:"PublishEventInfo"` // PUBLISHJOBFINISH

	// --------群发消息--------
	TotalCount           int `xml:"TotalCount"`
	FilterCount          int `xml:"FilterCount"`
	SentCount            int `xml:"SentCount"`
	ErrorCount           int `xml:"ErrorCount"`
	CopyrightCheckResult *struct {
		Count      int             `xml:"Count"`
		CheckState int             `xml:"CheckState"`
		ResultList []MassJobResult `xml:"ResultList>item"`
	} `xml:"CopyrightCheckResult"`
	ArticleUrlResult *struct {
		Count      int `xml:"Count"`
		ResultList []struct {
			ArticleIdx int    `xml:"ArticleIdx"`
			ArticleUrl string `xml:"ArticleUrl"`
		} `xml:"ResultList>item"`
	} `xml:"ArticleUrlResult"`

	// --------卡券消息--------
}
