package gateway

import (
	"context"

	"google.golang.org/api/youtube/v3"
)

type YouTubeGateway interface {
	FetchPlaylists(ctx context.Context) ([]*youtube.Playlist, error)
	FetchVideos(ctx context.Context, playlistId string) ([]*youtube.PlaylistItem, error)
}
