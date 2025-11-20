package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"youtube-playlist-video-tracker/src/infrastructure"
	"youtube-playlist-video-tracker/src/infrastructure/jsonstore"
	"youtube-playlist-video-tracker/src/usecase"
	"youtube-playlist-video-tracker/src/usecase/converter"
	"youtube-playlist-video-tracker/src/usecase/gateway"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const (
	credentialsFilePath string = "config/credentials.json"
	playlistsFilePath   string = "../playlists.json"
)

func main() {
	bootTime := time.Now().Format("2006-01-02_15-04-05")
	ctx := context.Background()

	credentials, err := loadCredentials(credentialsFilePath)
	if err != nil {
		fmt.Println("Error loading credentials:", err)
		return
	}

	config, err := google.ConfigFromJSON(credentials, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("OAuth設定の読み込み失敗: %v", err)
		return
	}

	var client gateway.YouTubeGateway
	clientImpl, err := infrastructure.NewYouTubeClient(ctx, getClient(config))
	if err != nil {
		log.Fatalf("YouTubeサービス作成失敗: %v", err)
		return
	}
	client = clientImpl

	playlistUc := usecase.NewPlaylistInteractor(client)
	currentPlaylists, err := playlistUc.BuildPlaylists(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(currentPlaylists) <= 0 {
		return
	}

	prevPlaylists, err := jsonstore.ReadPlaylistsFromJson(playlistsFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(prevPlaylists) > 0 {
		unavailableVideoUc := usecase.NewUnavailableVideoInteractor()
		unavailableVideos := unavailableVideoUc.DetectUnavailableVideos(converter.ToPlaylistEntities(prevPlaylists), currentPlaylists)

		if len(unavailableVideos) > 0 {
			if err := jsonstore.WriteUnavailableVideosToJson("../Unavailable Videos/"+bootTime+".json", converter.ToUnavailableVideoDTOs(unavailableVideos)); err != nil {
				return
			}

		}
	}

	jsonstore.WritePlaylistsToJson(playlistsFilePath, converter.ToPlaylistDTOs(currentPlaylists))
}

func loadCredentials(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func getClient(config *oauth2.Config) *http.Client {
	tokenFile := "config/token.json"
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("以下のURLをブラウザで開いて認証してください:\n%v\n", authURL)

	fmt.Print("認可コードを入力: ")
	var code string
	fmt.Scan(&code)

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("トークン交換失敗: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("トークン保存失敗: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
