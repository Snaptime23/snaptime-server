package cache

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/rpc_pb/videoApi"
)

type redisVideo struct {
	Id            string `json:"id" redis:"id"`
	UserId        string `json:"user_id" redis:"user_id"`
	PlayUrl       string `json:"play_url" redis:"play_url"`
	CoverUrl      string `json:"cover_url" redis:"cover_url"`
	FavoriteCount int64  `json:"favorite_count" redis:"favorite_count"`
	CommentCount  int64  `json:"comment_count" redis:"comment_count"`
	Title         string `json:"title" redis:"title"`
}

func VideoIsExists(ctx context.Context, vids ...string) int64 {
	keys := make([]string, len(vids))
	for i, vid := range vids {
		keys[i] = GetVideoMsgKey(vid)
	}
	return Exists(ctx, keys...)
}

func SetVideoMessage(ctx context.Context, video *videoApi.VideoInfo) (ok bool) {
	return HSet(ctx, GetVideoMsgKey(video.VideoID), &redisVideo{
		Id:            video.VideoID,
		UserId:        video.Author.UserId,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	})
}

func IncrVideoField(ctx context.Context, vid, field string, incr int64) (ok bool) {
	return HIncr(ctx, GetVideoMsgKey(vid), field, incr)
}

func GetVideoFields(ctx context.Context, vid string, field ...string) []interface{} {
	return HMGet(ctx, GetVideoMsgKey(vid), field...)
}

func GetVideoMessage(ctx context.Context, vid string) (*videoApi.VideoInfo, error) {
	key := GetVideoMsgKey(vid)
	value := HGetAll(ctx, key)
	video := new(videoApi.VideoInfo)
	video.VideoID = value["user_id"]
	video.Title = value["title"]
	video.FavoriteCount = GetNum(value, "favorite_count")
	video.PlayUrl = value["play_url"]
	video.CommentCount = GetNum(value, "comment_count")
	video.CoverUrl = value["cover_url"]
	return video, nil
}
