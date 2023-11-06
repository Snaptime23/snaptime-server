package cache

import (
	"context"
	"strconv"
)

func LikeIsExists(ctx context.Context, uids ...string) int64 {
	keys := make([]string, len(uids))
	for i, uid := range uids {
		keys[i] = GetUserLikeListKey(uid)
	}
	return Exists(ctx, keys...)
}

func IsLike(ctx context.Context, uid, vid string) bool {
	res := HMGet(ctx, GetUserLikeListKey(uid), vid)
	if res[0] == nil || res[0].(string) == "0" {
		return false
	}
	return true
}

func SetFavoriteList(ctx context.Context, userID string, kv ...string) bool {
	return HSet(ctx, GetUserLikeListKey(userID), kv)
}

func FavoriteAction(ctx context.Context, uid, vid string, action int64) bool {
	return HIncr(ctx, GetUserLikeListKey(uid), vid, action)
}

func GetAllUserLikes(ctx context.Context, uid string) (userLikes []string) {
	userLikes = make([]string, 0)
	res := HGetAll(ctx, GetUserLikeListKey(uid))
	for k, v := range res {
		vid := k
		action, _ := strconv.ParseInt(v, 10, 64)
		if action == 1 {
			userLikes = append(userLikes, vid)
		}
	}
	return
}

func GetFavoriteList(ctx context.Context, userID string) map[string]string {
	return HGetAll(ctx, GetUserLikeListKey(userID))
}
