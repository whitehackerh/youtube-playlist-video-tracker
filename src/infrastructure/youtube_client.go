package infrastructure

import (
	"context"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTubeClient struct {
	service *youtube.Service
}

func NewYouTubeClient(ctx context.Context, client *http.Client) (*YouTubeClient, error) {
	svc, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return &YouTubeClient{service: svc}, nil
}

func (c *YouTubeClient) FetchPlaylists(ctx context.Context) ([]*youtube.Playlist, error) {
	var result []*youtube.Playlist
	pageToken := ""

	for {
		resp, err := c.service.Playlists.List([]string{"snippet", "status", "contentDetails"}).Mine(true).MaxResults(50).PageToken(pageToken).Do()
		if err != nil {
			return nil, err
		}
		result = append(result, resp.Items...)
		// NextPageTokenには1度のコールで取得できる最大件数を上回って要素が存在する場合に提供される
		// NextPageTokenをPageTokenの引数に渡すことで次のデータを取得できる
		if resp.NextPageToken == "" {
			break
		}
		pageToken = resp.NextPageToken
	}

	return result, nil
}
