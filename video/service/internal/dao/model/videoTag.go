package model

import (
	"database/sql"
	"time"
)

const TableNameVideoTag = "video_tag"

type VideoTag struct {
	TagId     string `gorm:"primarykey"`
	TagName   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	VideoId   string       `gorm:"index"`
}

// TableName Comment's table name
func (*VideoTag) TableName() string {
	return TableNameVideoTag
}
