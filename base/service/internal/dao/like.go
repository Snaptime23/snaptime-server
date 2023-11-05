package dao

import (
	"context"
	"errors"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserLikeRecords(ctx context.Context, uid string) ([]string, error) {
	var likeVideoIds []string
	err := DB.WithContext(ctx).Model(&model.Like{}).
		Select("video_id").
		Where("user_id = ? AND action=0", uid).
		Find(&likeVideoIds).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return likeVideoIds, nil
		}
		return nil, err
	}
	return likeVideoIds, nil
}

func HasLiked(ctx context.Context, uid, vid string) (bool, error) {
	var ret bool
	err := DB.WithContext(ctx).Model(model.Like{}).
		Select("action").
		Where("user_id=? AND video_id=?", uid, vid).
		First(&ret).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return ret, nil
}

func UpdateAndInsertLikeRecord(ctx context.Context, uid, vid string, action int64) error {
	record := &model.Like{
		LikeID:  uuid.NewString(),
		UserId:  uid,
		VideoId: vid,
		Action:  action,
	}
	history := &model.Like{}
	err := DB.WithContext(ctx).Model(model.Like{}).
		Where("user_id=? AND video_id=?", uid, vid).
		First(history).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			err = DB.WithContext(ctx).Model(model.Like{}).
				Create(record).Error
			if err != nil {
				return err
			}
			//err = DB.WithContext(ctx).Model(model.Comment{}).
			//	Where("video_id = ?", vid).
			//	UpdateColumn("like_count", gorm.Expr("like_count+1")).Error
		}
		return err
	}
	if history.LikeID != "" && history.Action != record.Action {
		err = DB.WithContext(ctx).Model(&model.Like{}).
			UpdateColumn("action_type", action).
			Error
		if err != nil {
			return err
		}
		//
		//err = DB.WithContext(ctx).Model(model.Comment{}).
		//	Where("video_id = ?", vid).
		//	UpdateColumn("like_count", gorm.Expr("like_count-1")).Error
	}
	return nil
}
