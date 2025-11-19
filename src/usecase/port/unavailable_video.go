package port

import (
	"youtube-playlist-video-tracker/src/entity"
)

type UnavalilableVideoUseCase interface {
	DetectUnavailableVideos(prev []entity.Playlist, current []entity.Playlist) []entity.UnavailableVideo
}
