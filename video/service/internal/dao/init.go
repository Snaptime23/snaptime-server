package dao

import (
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Init() {
	if DB != nil {
		return
	}
	password := os.Getenv("DB_PASSWORD")
	dsn := os.Getenv("MYSQL_USERNAME") +
		":" +
		password +
		"@tcp(" +
		os.Getenv("DB_ADDR") +
		")/snaptime?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(
		&model.Video{},
		&model.Definition{},
		&model.VideoTag{},
	)
	if err != nil {
		panic(err)
	}
}
