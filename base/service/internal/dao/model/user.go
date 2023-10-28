package model

import (
	"database/sql"
	"time"
)

const TableNameUser = "user"

type User struct {
	ID              int64 `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime `gorm:"index"`
	UserName        string
	Password        string
	Avatar          string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        int64
	PublishNum      int64
	FavouriteNum    int64
	LikeNum         int64
	ReceivedLikeNum int64
}

// TableName Comment's table name
func (*User) TableName() string {
	return TableNameUser
}
