package converter

import (
	"youtube-playlist-video-tracker/src/entity"
	"youtube-playlist-video-tracker/src/infrastructure/jsonstore"
)

func ToVideoDTO(v entity.Video) jsonstore.Video {
	return jsonstore.Video{
		Id:           v.Id(),
		ChannelId:    v.ChannelId(),
		ChannelTitle: v.ChannelTitle(),
		Title:        v.Title(),
	}
}

func ToVideoDTOs(videos []entity.Video) []jsonstore.Video {
	var dtos []jsonstore.Video
	for _, v := range videos {
		dtos = append(dtos, ToVideoDTO(v))
	}
	return dtos
}
