package cache

import (
	"fmt"
	"strconv"
)

func GetUserInfoKey(userId string) string {
	return fmt.Sprintf("user_info_%s", userId)
}

func GetUserFollowListKey(userId int64) string {
	return fmt.Sprintf("user_follow_list_%d", userId)
}

func GetUserFollowerListKey(userID int64) string {
	return fmt.Sprintf("user_follower_list_%d", userID)
}

func GetUserLikeListKey(userID string) string {
	return fmt.Sprintf("user_like_list_%s", userID)
}

func GetVideoMsgKey(videoID string) string {
	return fmt.Sprintf("video_message_%d", videoID)
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
