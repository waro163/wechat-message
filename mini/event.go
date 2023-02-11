package mini

// EventType 事件类型
type EventType string

const (
	// EventWxaMediaCheck 异步校验图片/音频是否含有违法违规内容推送事件
	EventWxaMediaCheck EventType = "wxa_media_check"
	// 订阅消息事件推送
	SubscribeMsgPopupEvent EventType = "subscribe_msg_popup_event"
	// 服务通知点击...管理消息
	SubscribeMsgChangeEvent EventType = "subscribe_msg_change_event"
	// 订阅消息发送结果
	SubscribeMsgSentEvent EventType = "subscribe_msg_sent_event"
	// 进入客服会话
	UserEnterTempsession EventType = "user_enter_tempsession"
)
