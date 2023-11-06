package cache

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"strconv"
)

type redisUser struct {
	Id            string `json:"id" redis:"id"`
	Name          string `json:"name" redis:"name"`
	FollowCount   int64  `json:"follow_count,omitempty" redis:"follow_count"`
	FollowerCount int64  `json:"follower_count,omitempty" redis:"follower_count"`
	WorkCount     int64  `json:"work_count,omitempty" redis:"work_count"`
	FavoriteCount int64  `json:"favorite_count,omitempty" redis:"favorite_count"`
}

func UserIsExists(ctx context.Context, uids ...string) int64 {
	keys := make([]string, len(uids))
	for i, val := range uids {
		keys[i] = GetUserInfoKey(val)
	}
	return Exists(ctx, keys...)
}

func GetNum(f map[string]string, s string) int64 {
	ret, _ := strconv.Atoi(f[s])
	return int64(ret)
}

func GetUserInfo(ctx context.Context, uid string) (*model.User, error) {
	key := GetUserInfoKey(uid)
	value := HGetAll(ctx, key)
	var user = new(model.User)
	user.UserName = value["name"]
	//user.UserID = GetNum(value, "id"
	user.FollowCount = GetNum(value, "follow_count")
	user.FollowerCount = GetNum(value, "follower_count")
	//user.FavoriteCount = GetNum(value, "favorite_count")
	user.VideoNum = GetNum(value, "video_num")
	return user, nil
}

func SetUserInfo(ctx context.Context, user *model.User) bool {
	key := GetUserInfoKey(user.UserID)
	return HSet(ctx, key, &redisUser{
		Id:            user.UserID,
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		WorkCount:     user.VideoNum,
		FavoriteCount: user.FavouriteNum,
	})
}

// 1. 根据key和字段名查找值，key或field不存在时，对应的值返回nil，需要调用方自己判断
// 2. 返回的类型都为string，调用方自行转换
func GetUserFields(ctx context.Context, userID string, field ...string) []interface{} {
	return HMGet(ctx, GetUserInfoKey(userID), field...)
}

func IncrUserField(ctx context.Context, userID, field string, incr int64) (ok bool) {
	return HIncr(ctx, GetUserInfoKey(userID), field, incr)
}
