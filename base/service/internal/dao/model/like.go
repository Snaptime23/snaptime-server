package model

import (
	"database/sql"
	"time"
)

type Like struct {
	LikeID    string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	UserId    string       `gorm:"column:user_id;uniqueIndex:idx_like_user_video"`
	VideoId   string       `gorm:"column:video_id;uniqueIndex:idx_like_user_video"`
	Action    int64        `gorm:"column:action" json:"action"`
}

func (Like) TableName() string {
	return "like"
}
