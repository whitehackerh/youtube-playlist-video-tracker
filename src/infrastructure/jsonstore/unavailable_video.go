package jsonstore

type UnavailableVideo struct {
	PlaylistId    string
	PlaylistTitle string
	VideoId       string
	ChannelId     string
	ChannelTitle  string
	VideoTitle    string
	Reason        string
	DetectedTime  string
}
