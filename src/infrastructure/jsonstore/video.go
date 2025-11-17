package jsonstore

type Video struct {
	Id           string `json:"id"`
	ChannelId    string `json:"channel_id"`
	ChannelTitle string `json:"channel_title"`
	Title        string `json:"title"`
}
