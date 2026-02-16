package port

import (
	"youtube-playlist-video-tracker/src/entity"
)

type UnavailableVideoUseCase interface {
	DetectUnavailableVideos(prev []entity.Playlist, current []entity.Playlist) []entity.UnavailableVideo
}
