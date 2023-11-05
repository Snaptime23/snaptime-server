package model

import "gorm.io/gorm"

const TableNameDefinition = "definition"

type Definition struct {
	gorm.Model
	VideoID     string
	ResourceKey string
	Type        string
	Quality     int64 // higher is better, 10 - 20 for video, 10 for audio
	//FrameRate   int64 // for video
	//BitRate int64
	//FrameWidth  int64  // for video
	//FrameHeight int64  // for video
	//Encoder string // aac / h264 / h265
	//SampleRate int64  // for audio
}

// TableName Comment's table name
func (*Definition) TableName() string {
	return TableNameDefinition
}
