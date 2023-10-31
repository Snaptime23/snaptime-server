package dao

import (
	"context"
	"errors"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"gorm.io/gorm"
)

func GetCommentList(ctx context.Context, vid string) ([]model.Comment, error) {
	res := make([]model.Comment, 0)
	err := DB.WithContext(ctx).
		Model(model.Comment{}).
		Where("video_id = ?", vid).Find(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}
	return res, nil
}

func CreateComment(ctx context.Context, comment *model.Comment) error {
	return DB.WithContext(ctx).Model(model.Comment{}).
		Create(comment).Error
}

func DeleteCommentByID(ctx context.Context, id string) error {
	return DB.WithContext(ctx).Model(&model.Comment{}).
		Where("id = ?", id).
		Delete(nil).Error
}
