package model

import (
	"database/sql"
	"time"
)

const TableNameVideo = "video"

type Video struct {
	VideoID        string `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
	VideoName      string       `gorm:"index"`
	CreateUserId   string
	Bio            string
	PlayUrl        string
	CoverUrl       string
	CommentCount   int64
	FavouriteCount int64
	UploadState    int64
	MetaState      int64
}

// TableName Comment's table name
func (*Video) TableName() string {
	return TableNameVideo
}
