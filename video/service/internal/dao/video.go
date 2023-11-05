package dao

import (
	"context"
	"fmt"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao/model"
)

func CreateVideo(ctx context.Context, video *model.Video) (err error) {
	tx := DB.Begin().WithContext(ctx)
	if err = tx.Create(video).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return
}

func GetVideosByUserId(ctx context.Context, uid int64) ([]model.Video, error) {
	var videos []model.Video
	fmt.Println(uid)
	err := DB.WithContext(ctx).Model(model.Video{}).Where("user_id = ?", uid).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func MGetByTime(ctx context.Context, latestTime int64) ([]model.Video, error) {
	tx := DB.WithContext(ctx).Model(model.Video{})
	var videos []model.Video
	err := tx.Order("created_at desc").Limit(30).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func GetVideoByVideoId(ctx context.Context, vid string) (*model.Video, error) {
	var video *model.Video
	err := DB.WithContext(ctx).Model(model.Video{}).Where("video_id = ?", vid).First(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

func UpdateVideo(ctx context.Context, vid int64, videoMap *map[string]interface{}) (err error) {
	err = DB.WithContext(ctx).Model(model.Video{}).Where("video_id = ?", vid).Updates(videoMap).Error
	if err != nil {
		return err
	}
	return nil
}
