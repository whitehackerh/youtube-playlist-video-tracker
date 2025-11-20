package converter

import (
	"youtube-playlist-video-tracker/src/entity"
	"youtube-playlist-video-tracker/src/infrastructure/jsonstore"
)

func ToUnavailableVideoDTO(u entity.UnavailableVideo) jsonstore.UnavailableVideo {
	return jsonstore.UnavailableVideo{
		PlaylistId:    u.PlaylistId(),
		PlaylistTitle: u.PlaylistTitle(),
		VideoId:       u.VideoId(),
		ChannelId:     u.ChannelId(),
		ChannelTitle:  u.ChannelTitle(),
		VideoTitle:    u.VideoTitle(),
		Reason:        u.Reason(),
		DetectedTime:  u.DetectedTime(),
	}
}

func ToUnavailableVideoDTOs(unavailableVideos []entity.UnavailableVideo) []jsonstore.UnavailableVideo {
	var dtos []jsonstore.UnavailableVideo
	for _, p := range unavailableVideos {
		dtos = append(dtos, ToUnavailableVideoDTO(p))
	}
	return dtos
}
