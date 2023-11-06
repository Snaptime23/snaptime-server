package cache

import (
	"context"
	"fmt"
	"strconv"
)

func IncrVideoField(ctx context.Context, videoId string, num int64) error {
	key := fmt.Sprintf("video:%s:likes", videoId)
	return Rs.IncrBy(ctx, key, num).Err()
}

func GetVideoLikesCount(videoId string) int64 {
	key := fmt.Sprintf("video:%s:likes", videoId)
	// 使用 GET 命令获取点赞数量并将其转换为整数
	val, _ := Rs.Get(context.Background(), key).Result()
	count, _ := strconv.ParseInt(val, 10, 64)
	return count
}
