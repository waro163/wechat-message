package message

import "github.com/waro163/wechat-message/utils"

// Text 文本消息
type Text struct {
	CommonMsg
	Content CDATA `xml:"Content"`
}

// NewText 初始化文本消息
func NewText(content string) *Text {
	text := new(Text)
	text.Content = CDATA(content)
	text.SetMsgType(MsgTypeText)
	text.SetCreateTime(utils.GetCurrTS())
	return text
}
