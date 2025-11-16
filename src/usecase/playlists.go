package usecase

import (
	"context"
	"fmt"
	"youtube-playlist-video-tracker/src/domain/entity"
	"youtube-playlist-video-tracker/src/usecase/gateway"
)

type PlaylistInteractor struct {
	client gateway.YouTubeGateway
}

func NewPlaylistInteractor(client gateway.YouTubeGateway) *PlaylistInteractor {
	return &PlaylistInteractor{client: client}
}

func (uc *PlaylistInteractor) BuildPlaylists(ctx context.Context) ([]entity.Playlist, error) {
	yPlaylists, err := uc.client.FetchPlaylists(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range yPlaylists {
		fmt.Printf("再生リスト: %s (%s) - %d本\n", item.Snippet.Title, item.Id, item.ContentDetails.ItemCount)
	}

	/* TODO
	再生リストのレスポンスだけでは動画の情報が足りないため、再生リストIDをもとに、再生リストに含まれる動画を取得するAPIをコール
	再生リスト1件ずつしか指定できない -> 再生リストの数分並行でリクエスト
	*/

	return []entity.Playlist{}, nil
}
