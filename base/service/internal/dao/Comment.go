package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getChildren(root string) []*model.Comment {
	res := make([]*model.Comment, 0)
	DB.Model(model.Comment{}).
		Where("root_id = ?", root). // 从来没写过这种不等于
		Find(&res)
	return res
}

func GetChildrenNum(vid, root string) int64 {
	var sum int64
	DB.Model(model.Comment{}).Where("video_id = ?", vid).Where("root_id = ?", root).Count(&sum)
	return sum
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

func GetPageQue(ctx context.Context, vid, rootID string, page int64) ([]*model.Comment, error) {
	res := make([]*model.Comment, 0)
	err := DB.WithContext(ctx).Model(model.Comment{}).
		Where("video_id = ?", vid).
		Where("root_id = ?", rootID).
		Limit(10).
		Offset(int((page - 1) * 10)).
		Find(&res).Error
	fmt.Println(res)
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
		Where("comment_id = ?", id).
		Delete(nil).Error
}

func UpdateCommentLikeCount(ctx context.Context, commentId, userId string, likeAction int64) error {
	if likeAction == 0 {
		err := DB.WithContext(ctx).
			Model(model.Comment{}).
			Where("comment_id = ?", commentId).
			UpdateColumn("like_count", gorm.Expr("like_count+1")).Error
		if err != nil {
			return err
		}
		err = DB.WithContext(ctx).
			Model(model.CommentLike{}).
			Create(&model.CommentLike{
				LikeID:    uuid.NewString(),
				UserId:    userId,
				CommentId: commentId,
				Action:    0,
			}).Error
		return err
	} else {
		err := DB.WithContext(ctx).Model(model.Comment{}).
			Where("comment_id = ?", commentId).
			UpdateColumn("like_count", gorm.Expr("like_count-1")).Error
		if err != nil {
			return err
		}

		err = DB.WithContext(ctx).
			Model(model.CommentLike{}).
			Where("comment_id = ? and user_id = ?", commentId, userId).
			Delete(nil).Error
		return err
	}
}

func HasLikeComment(ctx context.Context, commentId, uid string) (int64, error) {
	like := &model.CommentLike{}
	err := DB.WithContext(ctx).
		Model(model.CommentLike{}).
		Where("comment_id = ? and user_id = ?", commentId, uid).
		First(&like).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return 0, nil
		}
	}

	return 1, nil
}
