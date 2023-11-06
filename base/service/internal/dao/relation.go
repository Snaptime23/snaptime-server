package dao

import (
	"context"
	"errors"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/Snaptime23/snaptime-server/v2/tools/errno"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFollowList(ctx context.Context, uid string) ([]string, error) {
	var userIds []string
	err := DB.WithContext(ctx).Model(model.Relation{}).
		Select("to_user_id").
		Where("user_id = ? AND action_type = 0", uid).
		Find(&userIds).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userIds, nil
		}
		return nil, err
	}
	return userIds, nil
}

func GetFollowerList(ctx context.Context, uid string) ([]string, error) {
	var userIds []string
	err := DB.WithContext(ctx).Model(model.Relation{}).
		Select("user_id").
		Where("to_user_id = ? AND action_type = 0", uid).
		Find(&userIds).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userIds, nil
		}
		return nil, err
	}
	return userIds, nil
}

func Follow(ctx context.Context, userId, toUserId string, action int64) error {
	histoty := &model.Relation{}
	err := DB.WithContext(ctx).
		Model(&model.Relation{}).
		Where("user_id = ? and to_user_id = ?", userId, toUserId).
		First(&histoty).Error
	if err == gorm.ErrRecordNotFound {
		if action == 1 {
			return errno.NewErrNo("不能取关没关注的人")
		}
		err = DB.WithContext(ctx).Model(model.User{}).
			Where("user_id = ?", toUserId).
			UpdateColumn("follower_count", gorm.Expr("follower_count+1")).Error
		err = DB.WithContext(ctx).Model(model.User{}).
			Where("user_id = ?", userId).
			UpdateColumn("follow_count", gorm.Expr("follow_count+1")).Error
		return DB.WithContext(ctx).Model(&model.Relation{}).Create(&model.Relation{
			RelationID: uuid.NewString(),
			UserId:     userId,
			ToUserId:   toUserId,
			ActionType: action,
		}).Error
	} else if err == nil {
		if action == histoty.ActionType {
			return errno.NewErrNo("重复操作")
		}
		if action == 1 {
			err = DB.WithContext(ctx).Model(model.User{}).
				Where("user_id = ?", toUserId).
				UpdateColumn("follower_count", gorm.Expr("follower_count-1")).Error
			err = DB.WithContext(ctx).Model(model.User{}).
				Where("user_id = ?", userId).
				UpdateColumn("follow_count", gorm.Expr("follow_count-1")).Error
		} else {
			err = DB.WithContext(ctx).Model(model.User{}).
				Where("user_id = ?", toUserId).
				UpdateColumn("follower_count", gorm.Expr("follower_count+1")).Error
			err = DB.WithContext(ctx).Model(model.User{}).
				Where("user_id = ?", userId).
				UpdateColumn("follow_count", gorm.Expr("follow_count+1")).Error
		}
		return DB.WithContext(ctx).
			Model(&model.Relation{}).
			Where("user_id = ? and to_user_id = ?", userId, toUserId).
			UpdateColumn("action_type", action).Error
	}
	return err
}

func IsFollowed(ctx context.Context, uid, toUserId int64) (bool, error) {
	var ret bool
	err := DB.WithContext(ctx).Model(model.Relation{}).
		Select("action").
		Where("user_id = ? AND to_user_id = ?", uid, toUserId).
		First(&ret).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return ret, nil
}

func UpdateAndInsertRelation(ctx context.Context, uid, toUserId string, action int64) error {
	relation := &model.Relation{
		UserId:     uid,
		ToUserId:   toUserId,
		ActionType: action,
	}
	err := DB.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "to_user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"action"}),
		}).Create(&relation).Error
	if err != nil {
		return err
	}
	return nil
}
