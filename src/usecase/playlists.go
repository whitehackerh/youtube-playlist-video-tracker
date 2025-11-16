package usecase

import (
	"context"
	"fmt"
	"sync"
	"youtube-playlist-video-tracker/src/domain/entity"
	"youtube-playlist-video-tracker/src/usecase/gateway"

	"google.golang.org/api/youtube/v3"
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

	// TODO DELETE
	for _, item := range yPlaylists {
		fmt.Printf("再生リスト: %s (%s) - %d本\n", item.Snippet.Title, item.Id, item.ContentDetails.ItemCount)
	}

	type result struct {
		playlist entity.Playlist
		err      error
	}

	results := make(chan result, len(yPlaylists))
	var wg sync.WaitGroup

	// NOTE
	// Playlistsのレスポンスだけでは動画の情報が足りないため、再生リストIDをもとに、再生リストに含まれる動画の情報を取得するAPI(PlayListItems)をコール
	// PlaylistItemsは再生リスト1つしか指定できない -> 再生リストの数分並行でコール
	for _, pl := range yPlaylists {
		wg.Add(1)
		go func(pl *youtube.Playlist) {
			defer wg.Done()

			items, err := uc.client.FetchVideos(ctx, pl.Id)
			if err != nil {
				results <- result{err: err}
				return
			}

			var videos []entity.Video
			for _, item := range items {
				if item.Snippet == nil || item.Snippet.ResourceId == nil {
					continue
				}
				videos = append(videos, entity.NewVideo(item.Snippet.ResourceId.VideoId, item.Snippet.Title, item.Snippet.VideoOwnerChannelId, item.Snippet.VideoOwnerChannelTitle))
			}

			results <- result{
				playlist: entity.NewPlaylist(pl.Id, pl.Snippet.Title, videos),
			}
		}(pl)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var playlists []entity.Playlist
	for r := range results {
		if r.err != nil {
			continue
		}
		playlists = append(playlists, r.playlist)
	}

	// TODO DELETE
	fmt.Println(playlists)

	return playlists, nil
}
