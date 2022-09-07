package message

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
	return voice
}
