package usecase

import (
	"time"
	"youtube-playlist-video-tracker/src/entity"
	"youtube-playlist-video-tracker/src/usecase/port"
)

type UnavailableVideoInteractor struct{}

func NewUnavailableVideoInteractor() port.UnavalilableVideoUseCase {
	return &UnavailableVideoInteractor{}
}

func (uc *UnavailableVideoInteractor) DetectUnavailableVideos(prev []entity.Playlist, current []entity.Playlist) []entity.UnavailableVideo {
	/*
		[NOTE]
			見れなくなった動画の定義
			・以前(前回システム実行時)は見れていたものの、今回システム実行した結果、非公開or削除されている動画
			・地域制限による視聴不可は除外
			・限定公開については再生リストに含まれていれば引き続き視聴可能のため除外

			非公開or削除された動画の特徴・・・playlistItems.listのレスポンスが以下になっている
			・チャンネルID, チャンネルタイトルが空文字
			・動画タイトルが"Private video" (非公開), "Deleted video" (削除)

			見れなくなった動画の判定は以下をすべて満たす場合とする
			・prevの動画のチャンネルIDが空文字ではない
			・prevの動画のチャンネルタイトルが空文字ではない
			・prevの動画タイトルが"Private video"ではない
			・prevの動画タイトルが"Deleted video"ではない
			・currentの動画のチャンネルIDが空文字
			・currentの動画のチャンネルタイトルが空文字
			・currentの動画タイトルが"Private video"または"Deleted video"
	*/

	currentVideos := make(map[string]entity.Video)
	for _, pl := range current {
		for _, v := range pl.Videos() {
			currentVideos[v.Id()] = v
		}
	}

	detectedTime := time.Now().Format(time.DateTime)
	var result []entity.UnavailableVideo
	for _, pl := range prev {
		for _, v := range pl.Videos() {
			c, ok := currentVideos[v.Id()]
			if !ok {
				// 再生リストから消えた → 対象外
				continue
			}
			if uc.isPreviouslyViewable(v) && uc.isNowUnavailable(c) {
				result = append(result, entity.NewUnavailableVideo(
					pl.Id(),
					pl.Title(),
					v.Id(),
					v.ChannelId(),
					v.ChannelTitle(),
					v.Title(),
					detectedTime,
					uc.classifyUnavailableReason(c.Title()),
				))
			}
		}
	}

	return result
}

func (uc *UnavailableVideoInteractor) isPreviouslyViewable(v entity.Video) bool {
	return v.ChannelId() != "" &&
		v.ChannelTitle() != "" &&
		v.Title() != "Private video" &&
		v.Title() != "Deleted video"
}

func (uc *UnavailableVideoInteractor) isNowUnavailable(v entity.Video) bool {
	return v.ChannelId() == "" &&
		v.ChannelTitle() == "" &&
		(v.Title() == "Private video" || v.Title() == "Deleted video")
}

func (uc *UnavailableVideoInteractor) classifyUnavailableReason(title string) string {
	switch title {
	case "Private video":
		return "Private"
	case "Deleted video":
		return "Deleted"
	default:
		return "Unknown"
	}
}
