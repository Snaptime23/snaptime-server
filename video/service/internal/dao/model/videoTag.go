package model

import (
	"gorm.io/gorm"
)

const TableNameVideoTag = "video_tag"

type VideoTag struct {
	gorm.Model
	TagName string `gorm:"index"`
	VideoId string `gorm:"index"`
}

// TableName Comment's table name
func (*VideoTag) TableName() string {
	return TableNameVideoTag
}
