package port

import (
	"context"
	"youtube-playlist-video-tracker/src/entity"
)

type PlaylistUseCase interface {
	BuildPlaylists(ctx context.Context) ([]entity.Playlist, error)
}
