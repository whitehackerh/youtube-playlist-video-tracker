package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	credentialsFilePath string = "config/credentials.json"
)

func main() {
	credentials, err := loadCredentials(credentialsFilePath)
	if err != nil {
		fmt.Println("Error loading credentials:", err)
		return
	}

	config, err := google.ConfigFromJSON(credentials, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("OAuth設定の読み込み失敗: %v", err)
	}

	client := getClient(config)

	// YouTubeサービスの作成
	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("YouTubeサービス作成失敗: %v", err)
	}

	// 自分の再生リストを取得
	call := service.Playlists.List([]string{"snippet", "status", "contentDetails"}).Mine(true).MaxResults(50)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("API呼び出し失敗: %v", err)
	}

	// 結果を表示
	for _, item := range response.Items {
		fmt.Printf("再生リスト: %s (%s) - %d本\n", item.Snippet.Title, item.Id, item.ContentDetails.ItemCount)
	}
	fmt.Println(response)
	// id : 再生リストID

	// JSONとして保存（任意）
	f, _ := os.Create("my_playlists.json")
	defer f.Close()
	json.NewEncoder(f).Encode(response)

	/* ----------------- */
	/* 試しに再生リストに含まれる動画リストを取得するAPIを1回コールしてみる
	// 再生リストのレスポンスだけでは動画の情報が足りないため、再生リストIDをもとに、再生リストに含まれる動画を取得するAPIをコール
	// 再生リスト1件ずつしか指定できない -> 再生リストの数分並行でリクエスト
	// 1度に最大50件の動画までしか取得できないため、再生リストに含まれる動画が51件以上ある場合は、複数回コール
	*/
	call2 := service.PlaylistItems.List([]string{"snippet"}).PlaylistId("test").MaxResults(50)
	response2, err := call2.Do()
	if err != nil {
		log.Fatalf("API呼び出し失敗: %v", err)
	}

	// JSONとして保存（任意）
	f2, _ := os.Create("my_playlist_items.json")
	defer f2.Close()
	json.NewEncoder(f2).Encode(response2)
	// items[n].resourceId.videoId : 動画ID
	// items[n].title : 動画タイトル
	// items[n].videoOwnerChannelId : 投稿者チャンネルID
	// items[n].videoOwnerChannelTitle : 投稿者チャンネルタイトル

	/* ----------------- */

	// 一致する動画IDの動画に動画の詳細をマージしていく

	// 旧情報と新情報を比較し、見れなくなった動画の情報を書き込む
}

func loadCredentials(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func getClient(config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
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
