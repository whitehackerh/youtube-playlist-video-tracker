package entity

type Video struct {
	id           string
	channelId    string
	channelTitle string
	title        string
}

func NewVideo(
	id string,
	channelId string,
	channelTitle string,
	title string,
) Video {
	return Video{
		id:           id,
		channelId:    channelId,
		channelTitle: channelTitle,
		title:        title,
	}
}

func (v *Video) Id() string {
	return v.id
}

func (v *Video) ChannelId() string {
	return v.channelId
}

func (v *Video) ChannelTitle() string {
	return v.channelTitle
}

func (v *Video) Title() string {
	return v.title
}
