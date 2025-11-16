package port

import (
	"context"
	"youtube-playlist-video-tracker/src/domain/entity"
)

type PlaylistUseCase interface {
	BuildPlaylists(ctx context.Context) ([]entity.Playlist, error)
}
