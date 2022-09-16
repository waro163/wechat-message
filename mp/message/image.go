package message

import "github.com/waro163/wechat-message/utils"

// Image 图片消息
type Image struct {
	CommonMsg
	MediaID CDATA `xml:"Image>MediaId"`

	// Image struct {
	// 	MediaID string `xml:"MediaId"`
	// } `xml:"Image"`
}

// NewImage 回复图片消息
func NewImage(mediaID string) *Image {
	image := new(Image)
	image.MediaID = CDATA(mediaID)
	image.SetMsgType(MsgTypeImage)
	image.SetCreateTime(utils.GetCurrTS())
	return image
}
