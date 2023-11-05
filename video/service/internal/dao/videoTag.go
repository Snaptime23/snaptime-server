package dao

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao/model"
	"github.com/google/uuid"
)

func CreateVideoTag(ctx context.Context, tagName, videoId string) error {
	return DB.WithContext(ctx).
		Model(&model.VideoTag{}).
		Create(&model.VideoTag{
			TagId:   uuid.NewString(),
			TagName: tagName,
			VideoId: videoId,
		}).
		Error
}

func SearchByVideoTag(ctx context.Context, tagName string) (videoId []string, err error) {
	videoId = make([]string, 0)
	err = DB.WithContext(ctx).
		Model(&model.VideoTag{}).
		Select("video_id").
		Where("tag_name = ?", tagName).
		Find(&videoId).
		Error
	if err != nil {
		return nil, err
	}
	return videoId, err
}
