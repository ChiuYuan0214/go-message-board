package entities

import "time"

type Image struct {
	UserId       uint64    `gorm:"column:user_id;primaryKey"`
	FileName     string    `gorm:"column:file_name"`
	Descript     string    `gorm:"column:descript"`
	CreationTime time.Time `gorm:"column:creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}

func (Image) TableName() string {
	return "images"
}
