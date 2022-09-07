package message

// Music 音乐消息
type Music struct {
	CommonMsg

	Music struct {
		Title        CDATA `xml:"Title,omitempty"`
		Description  CDATA `xml:"Description,omitempty"`
		MusicURL     CDATA `xml:"MusicUrl,omitempty"`
		HQMusicURL   CDATA `xml:"HQMusicUrl,omitempty"`
		ThumbMediaID CDATA `xml:"ThumbMediaId"`
	} `xml:"Music"`
}

// NewMusic  回复音乐消息
func NewMusic(title, description, musicURL, hQMusicURL, thumbMediaID string) *Music {
	music := new(Music)
	music.Music.Title = CDATA(title)
	music.Music.Description = CDATA(description)
	music.Music.MusicURL = CDATA(musicURL)
	music.Music.ThumbMediaID = CDATA(thumbMediaID)
	return music
}
