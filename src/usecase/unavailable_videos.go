package usecase

import (
	"context"
	"youtube-playlist-video-tracker/src/entity"
	"youtube-playlist-video-tracker/src/usecase/gateway"
	"youtube-playlist-video-tracker/src/usecase/port"
)

type UnavailableVideoInteractor struct {
	client gateway.YouTubeGateway
}

func NewUnavailableVideoInteractor(client gateway.YouTubeGateway) port.UnavalilableVideoUseCase {
	return &UnavailableVideoInteractor{client: client}
}

func (uc *UnavailableVideoInteractor) DetectUnavailableVideos(ctx context.Context, prev []entity.Playlist, current []entity.Playlist) ([]entity.UnavailableVideo, error) {

	type snapshots struct {
		playlistId    string
		playlistTitle string
		videoId       string
		channelId     string
		channelTitle  string
		videoTitle    string
	}

	return []entity.UnavailableVideo{}, nil
}
