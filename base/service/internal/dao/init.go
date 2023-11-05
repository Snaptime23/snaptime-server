package dao

import (
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	if DB != nil {
		return
	}
	dsn := "jiyeon:998244353@tcp(d.reeky.org:13307)/snaptime?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(
		&model.User{},
		&model.Comment{},
		&model.Like{},
		&model.Relation{},
	)
	if err != nil {
		panic(err)
	}
}
