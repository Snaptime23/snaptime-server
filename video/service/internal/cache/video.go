package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type redisVideo struct {
	Id            int64  `json:"id" redis:"id"`
	UserId        int64  `json:"user_id" redis:"user_id"`
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

//func SetVideoMessage(ctx context.Context, video *model.Video) (ok bool) {
//	return HSet(ctx, GetVideoMsgKey(int64(video.ID)), &redisVideo{
//		Id:            int64(video.ID),
//		UserId:        video.UserId,
//		PlayUrl:       video.PlayUrl,
//		CoverUrl:      video.CoverUrl,
//		FavoriteCount: video.FavoriteCount,
//		CommentCount:  video.CommentCount,
//		Title:         video.Title,
//	})
//}

func IncrVideoField(ctx context.Context, vid, field string, incr int64) (ok bool) {
	return HIncr(ctx, GetVideoMsgKey(vid), field, incr)
}

func GetVideoFields(ctx context.Context, vid string, field ...string) []interface{} {
	return HMGet(ctx, GetVideoMsgKey(vid), field...)
}

//func GetVideoMessage(ctx context.Context, vid int64) (*model.Video, error) {
//	key := GetVideoMsgKey(vid)
//	value := HGetAll(ctx, key)
//	video := new(model.Video)
//	video.UserId = GetNum(value, "user_id")
//	video.Title = value["title"]
//	video.FavoriteCount = GetNum(value, "favorite_count")
//	video.PlayUrl = value["play_url"]
//	video.CommentCount = GetNum(value, "comment_count")
//	video.CoverUrl = value["cover_url"]
//	return video, nil
//}

// ========通用操作==========

func Exists(ctx context.Context, key ...string) int64 {
	return Rs.Exists(ctx, key...).Val()
}

// =======String操作==============

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (ok bool) {
	if value == nil {
		return false
	}
	bs, err := json.Marshal(value)
	if err != nil {
		//Log.Warnf("marshal value err: %v", err)
		return false
	}
	if expiration == 0 {
		expiration = time.Hour
	}
	if status := Rs.Set(ctx, key, bs, expiration); status.Err() != nil {
		//Log.Errorf("set %v to redis err: %v", key, status.Err())
		return false
	}
	return true
}

func Incr(ctx context.Context, key string) (cnt int64, err error) {
	cnt, err = Rs.Incr(ctx, key).Result()
	if err != nil {
		//Log.Warnf("Incr %v redis err: %v", key, err)
		return
	}
	return
}

// ==========Hash操作============

func HSet(ctx context.Context, key string, value interface{}) (ok bool) {
	if status := Rs.HSet(ctx, key, value); status.Err() != nil {
		return false
	}
	return true
}

// HGetAll 返回值为map，需要调用方自行转换
func HGetAll(ctx context.Context, key string) map[string]string {
	return Rs.HGetAll(ctx, key).Val()
}

func HMGet(ctx context.Context, key string, field ...string) []interface{} {
	return Rs.HMGet(ctx, key, field...).Val()
}

func HIncr(ctx context.Context, key, field string, incr int64) (ok bool) {
	if status := Rs.HIncrBy(ctx, key, field, incr); status.Err() != nil {
		return false
	}
	return true
}

func GetUserInfoKey(userId int64) string {
	return fmt.Sprintf("user_info_%d", userId)
}

func GetUserFollowListKey(userId int64) string {
	return fmt.Sprintf("user_follow_list_%d", userId)
}

func GetUserFollowerListKey(userID int64) string {
	return fmt.Sprintf("user_follower_list_%d", userID)
}

func GetUserLikeListKey(userID int64) string {
	return fmt.Sprintf("user_like_list_%d", userID)
}

func GetVideoMsgKey(videoID string) string {
	return fmt.Sprintf("video_message_%s", videoID)
}

func GetIDFromUserMsgKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[10:], 10, 64)
	return
}

func GetIDFromUserLikeListKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[15:], 10, 64)
	return
}

func GetIDFromUserFollowListKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[17:], 10, 64)
	return
}

func GetIDFromVideoMsgKey(key string) (id int64) {
	id, _ = strconv.ParseInt(key[14:], 10, 64)
	return
}

func GetFavoriteLmtKey(ipaddr string) string {
	return "favorite_limit_" + ipaddr
}
