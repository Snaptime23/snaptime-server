package model

import (
	"database/sql"
	"time"
)

const TableNameComment = "comment"

type Comment struct {
	CommentID   string `gorm:"primarykey"`
	RootID      string
	ParentID    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
	LikeCount   int64
	UserId      string `json:"user_id" gorm:"index:idx_comment_user_id"`
	UserName    string
	UserAvatar  string
	VideoId     string     `json:"video_id"`
	Content     string     `json:"content"`
	PublishDate int64      `json:"publish_date"`
	Children    []*Comment `gorm:"-" json:"children"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
