package message

import "github.com/waro163/wechat-message/utils"

// Voice 语音消息
type Voice struct {
	CommonMsg
	MediaID CDATA `xml:"Voice>MediaId"`
	// Voice struct {
	// 	MediaID string `xml:"MediaId"`
	// } `xml:"Voice"`
}

// NewVoice 回复语音消息
func NewVoice(mediaID string) *Voice {
	voice := new(Voice)
	voice.MediaID = CDATA(mediaID)
	voice.SetMsgType(MsgTypeVoice)
	voice.SetCreateTime(utils.GetCurrTS())
	return voice
}
