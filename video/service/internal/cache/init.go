package cache

import (
	"context"
	"fmt"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao"
	"github.com/go-co-op/gocron"
	"time"

	//"github.com/go-co-op/gocron"
	"github.com/redis/go-redis/v9"
	"os"
	//"time"
)

var Rs *redis.Client

func InitRedis() {
	if Rs != nil {
		return
	}
	Rs = redis.NewClient(&redis.Options{
		Addr:     "d.reeky.org:32771",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := Rs.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	s := gocron.NewScheduler(time.Local)
	_, _ = s.Every(5).Second().Do(func() {
		go func() {
			// 获取所有视频的点赞数量
			videoIDs, _ := Rs.Keys(context.Background(), "video:*:likes").Result()

			for _, key := range videoIDs {
				// 获取视频点赞数量
				count, _ := Rs.Get(context.Background(), key).Int64()

				// 解析 videoID
				var videoID string
				_, err = fmt.Sscanf(key, "video:%s:likes", &videoID)

				// 使用 dao.UpdateVideo 函数将点赞数量更新到 MySQL
				_ = dao.UpdateVideo(context.Background(), videoID, &map[string]interface{}{"favorite_count": count})
			}
		}()
	})
	//s.StartBlocking()
}
