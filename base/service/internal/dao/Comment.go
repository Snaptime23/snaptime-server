package dao

import (
	"context"
	"errors"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"gorm.io/gorm"
)

func getChildren(root string) []*model.Comment {
	res := make([]*model.Comment, 0)
	DB.Model(model.Comment{}).
		Where("root_id = ?", root). // 从来没写过这种不等于
		Find(res)
	return res
}

func GetCommentCount(vid string) int64 {
	var sum int64
	if err := DB.Model(model.Comment{}).
		Where("video_id = ?", vid).
		Where("root_id = ?", "").
		Count(&sum).Error; err != nil {
		return 0
	}
	return sum
}

func GetPageQue(ctx context.Context, vid string, page int64) ([]*model.Comment, error) {
	res := make([]*model.Comment, 0)
	err := DB.WithContext(ctx).Model(model.Comment{}).
		Where("video_id = ?", vid).
		Limit(10).
		Offset(int((page - 1) * 10)).
		Find(res).Error
	for i, _ := range res {
		res[i].Children = getChildren(res[i].CommentID)
	}
	return res, err
}

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
	for i, _ := range res {
		res[i].Children = getChildren(res[i].CommentID)
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
