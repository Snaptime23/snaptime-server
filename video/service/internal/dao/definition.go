package dao

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao/model"
)

func CreateVideoDefinition(ctx context.Context, definition *model.Definition) (err error) {
	tx := DB.Begin().WithContext(ctx)
	if err = tx.Create(definition).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return
}

func GetHighUrl(ctx context.Context, videoId string) (string, error) {
	ret := ""
	err := DB.WithContext(ctx).
		Model(&model.Definition{}).
		Select("resource_key").
		Where("video_id = ? and quality = ?", videoId, 20).
		Find(&ret).
		Error
	return ret, err
}
