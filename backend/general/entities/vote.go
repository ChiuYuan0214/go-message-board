package entities

import "time"

type Vote struct {
	VoteId       uint64    `gorm:"primaryKey"`
	UserId       uint64    `gorm:"column:user_id"`
	SourceId     uint64    `gorm:"column:source_id"`
	Score        int8      `gorm:"column:score"`
	VoteType     string    `gorm:"column:vote_type"`
	CreationTime time.Time `gorm:"column:creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}
