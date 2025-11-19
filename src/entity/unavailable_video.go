package entity

type UnavailableVideo struct {
	playlistId    string
	playlistTitle string
	videoId       string
	channelId     string
	channelTitle  string
	videoTitle    string
	reason        string
	detectedTime  string
}

func NewUnavailableVideo(
	playlistId string,
	playlistTitle string,
	videoId string,
	channelId string,
	channelTitle string,
	videoTitle string,
	reason string,
	detectedTime string,
) UnavailableVideo {
	return UnavailableVideo{
		playlistId:    playlistId,
		playlistTitle: playlistTitle,
		videoId:       videoId,
		channelId:     channelId,
		channelTitle:  channelTitle,
		videoTitle:    videoTitle,
		reason:        reason,
		detectedTime:  detectedTime,
	}
}

func (u *UnavailableVideo) PlaylistId() string {
	return u.playlistId
}

func (u *UnavailableVideo) PlaylistTitle() string {
	return u.playlistTitle
}

func (u *UnavailableVideo) VideoId() string {
	return u.videoId
}

func (u *UnavailableVideo) ChannelId() string {
	return u.channelId
}

func (u *UnavailableVideo) ChannelTitle() string {
	return u.channelTitle
}

func (u *UnavailableVideo) VideoTitle() string {
	return u.videoTitle
}

func (u *UnavailableVideo) Reason() string {
	return u.reason
}

func (u *UnavailableVideo) DetectedTime() string {
	return u.detectedTime
}
