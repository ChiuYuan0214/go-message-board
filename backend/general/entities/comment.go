package entities

import "time"

type Comment struct {
	CommentId  uint64    `gorm:"primaryKey"`
	UserId     uint64    `gorm:"column:userId"`
	ArticleId  uint64    `gorm:"column:article_id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	Edited     bool      `gorm:"column:edited"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
