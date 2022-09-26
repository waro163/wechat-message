package mp

// EventType 事件类型
type EventType string

const (
	// --------基础事件推送--------

	// EventSubscribe 订阅
	EventSubscribe EventType = "subscribe"
	// EventUnsubscribe 取消订阅
	EventUnsubscribe EventType = "unsubscribe"
	// EventScan 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventScan EventType = "SCAN"
	// EventLocation 上报地理位置事件
	EventLocation EventType = "LOCATION"
	// EventClick 点击菜单拉取消息时的事件推送
	EventClick EventType = "CLICK"
	// EventView 点击菜单跳转链接时的事件推送
	EventView EventType = "VIEW"

	// --------自定义菜单事件推送--------

	// EventScancodePush 扫码推事件的事件推送
	EventScancodePush EventType = "scancode_push"
	// EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventScancodeWaitmsg EventType = "scancode_waitmsg"
	// EventPicSysphoto 弹出系统拍照发图的事件推送
	EventPicSysphoto EventType = "pic_sysphoto"
	// EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
	EventPicPhotoOrAlbum EventType = "pic_photo_or_album"
	// EventPicWeixin 弹出微信相册发图器的事件推送
	EventPicWeixin EventType = "pic_weixin"
	// EventLocationSelect 弹出地理位置选择器的事件推送
	EventLocationSelect EventType = "location_select"
	// EventViewMiniprogram 点击菜单跳转小程序的事件推送
	EventViewMiniprogram EventType = "view_miniprogram"

	//  --------模版消息事件推送--------

	// EventTemplateSendJobFinish 发送模板消息推送通知
	EventTemplateSendJobFinish EventType = "TEMPLATESENDJOBFINISH"

	// --------群发消息事件推送--------

	// EventMassSendJobFinish 群发消息推送通知
	EventMassSendJobFinish EventType = "MASSSENDJOBFINISH"

	//

	// EventWxaMediaCheck 异步校验图片/音频是否含有违法违规内容推送事件
	EventWxaMediaCheck EventType = "wxa_media_check"

	// --------订阅通知事件推送--------

	// EventSubscribeMsgPopupEvent 订阅通知事件推送（图文场景）
	EventSubscribeMsgPopupEvent EventType = "subscribe_msg_popup_event"
	// EventSubscribeMsgChangeEvent 订阅通知修改事件
	EventSubscribeMsgChangeEvent EventType = "subscribe_msg_change_event"
	// EventSubscribeMsgSendEvent 订阅通知事件推送（bizsend接口发送）
	EventSubscribeMsgSendEvent EventType = "subscribe_msg_sent_event"

	// --------发布能力事件推送--------

	// EventPublishJobFinish 发布任务完成
	EventPublishJobFinish EventType = "PUBLISHJOBFINISH"
)
