package model

import (
	"database/sql"
	"time"
)

type CommentLike struct {
	LikeID    string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	UserId    string       `gorm:"column:user_id;uniqueIndex:idx_like_user_comment"`
	CommentId string       `gorm:"column:comment_id;uniqueIndex:idx_like_user_comment"`
	Action    int64        `gorm:"column:action" json:"action"`
}

func (CommentLike) TableName() string {
	return "comment_like"
}
