package entity

type PlayList struct {
	id     string
	title  string
	videos []Video
}

func NewPlayList(
	id string,
	title string,
	videos []Video,
) PlayList {
	return PlayList{
		id:     id,
		title:  title,
		videos: videos,
	}
}

func (p *PlayList) Id() string {
	return p.id
}

func (p *PlayList) Title() string {
	return p.title
}

func (p *PlayList) Videos() []Video {
	return p.videos
}
