package port

import (
	"context"
	"youtube-playlist-video-tracker/src/entity"
)

type UnavalilableVideoUseCase interface {
	DetectUnavailableVideos(ctx context.Context, prev []entity.Playlist, current []entity.Playlist) ([]entity.UnavailableVideo, error)
}
