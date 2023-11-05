package model

import (
	"database/sql"
	"time"
)

type Relation struct {
	RelationID string `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	UserId     string       `gorm:"column:user_id;not null;uniqueIndex:idx_relation"`
	ToUserId   string       `gorm:"column:to_user_id;not null;uniqueIndex:idx_relation"`
	Action     int64        `gorm:"column:action" json:"action"`
}
