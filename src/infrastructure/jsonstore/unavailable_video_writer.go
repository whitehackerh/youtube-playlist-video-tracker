package jsonstore

import (
	"encoding/json"
	"os"
)

func WriteUnavailableVideosToJson(path string, unavailableVideos []UnavailableVideo) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	return encoder.Encode(unavailableVideos)
}
