package jsonstore

import (
	"encoding/json"
	"os"
)

func ReadPlaylistsFromJson(path string) ([]Playlist, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []Playlist{}, err
	}

	var playlist []Playlist
	if err := json.Unmarshal(data, &playlist); err != nil {
		return []Playlist{}, err
	}

	return playlist, nil
}
