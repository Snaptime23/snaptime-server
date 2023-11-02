package model

import (
	"database/sql"
	"time"
)

const TableNameUser = "video"

type Video struct {
	VideoID   string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	VideoName string       `gorm:"index"`
	Bio       string
}

// TableName Comment's table name
func (*Video) TableName() string {
	return TableNameUser
}
