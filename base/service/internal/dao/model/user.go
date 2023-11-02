package model

import (
	"database/sql"
	"time"
)

const TableNameUser = "user"

type User struct {
	UserID        string `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
	UserName      string       `gorm:"index"`
	Password      string
	Bio           string
	Avatar        string
	FollowCount   int64
	FollowerCount int64
	VideoNum      int64
	FavouriteNum  int64
}

// TableName Comment's table name
func (*User) TableName() string {
	return TableNameUser
}
