package converter

import (
	"youtube-playlist-video-tracker/src/entity"
	"youtube-playlist-video-tracker/src/infrastructure/jsonstore"
)

func ToPlaylistDTO(p entity.Playlist) jsonstore.Playlist {
	return jsonstore.Playlist{
		Id:     p.Id(),
		Title:  p.Title(),
		Videos: ToVideoDTOs(p.Videos()),
	}
}

func ToPlaylistDTOs(playlists []entity.Playlist) []jsonstore.Playlist {
	var dtos []jsonstore.Playlist
	for _, p := range playlists {
		dtos = append(dtos, ToPlaylistDTO(p))
	}
	return dtos
}

func ToPlaylistEntity(p jsonstore.Playlist) entity.Playlist {
	return entity.NewPlaylist(p.Id, p.Title, ToVideoEntities(p.Videos))
}

func ToPlaylistEntities(playlists []jsonstore.Playlist) []entity.Playlist {
	var entities []entity.Playlist
	for _, p := range playlists {
		entities = append(entities, ToPlaylistEntity(p))
	}
	return entities
}
