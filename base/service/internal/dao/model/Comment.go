package model

import (
	"database/sql"
	"time"
)

const TableNameComment = "comment"

type Comment struct {
	CommentID   string `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
	UserId      string       `json:"user_id" gorm:"index:idx_comment_user_id"`
	VideoId     string       `json:"video_id"`
	Content     string       `json:"content"`
	PublishDate string       `json:"publish_date"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
