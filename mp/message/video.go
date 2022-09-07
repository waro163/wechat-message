package message

// Video 视频消息
type Video struct {
	CommonMsg

	Video struct {
		MediaID     CDATA `xml:"MediaId"`
		Title       CDATA `xml:"Title,omitempty"`
		Description CDATA `xml:"Description,omitempty"`
	} `xml:"Video"`
}

// NewVideo 回复图片消息
func NewVideo(mediaID, title, description string) *Video {
	video := new(Video)
	video.Video.MediaID = CDATA(mediaID)
	video.Video.Title = CDATA(title)
	video.Video.Description = CDATA(description)
	return video
}
