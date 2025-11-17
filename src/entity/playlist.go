package entity

type Playlist struct {
	id     string
	title  string
	videos []Video
}

func NewPlaylist(
	id string,
	title string,
	videos []Video,
) Playlist {
	return Playlist{
		id:     id,
		title:  title,
		videos: videos,
	}
}

func (p *Playlist) Id() string {
	return p.id
}

func (p *Playlist) Title() string {
	return p.title
}

func (p *Playlist) Videos() []Video {
	return p.videos
}
