package model

const TableNameDefinition = "definition"

type Definition struct {
	VideoID string `gorm:"primarykey"`
	One     string
	Two     string
	Three   string
	Four    string
	Five    string
}

// TableName Comment's table name
func (*Definition) TableName() string {
	return TableNameDefinition
}
