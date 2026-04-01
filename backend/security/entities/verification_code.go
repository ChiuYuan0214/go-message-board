package entities

import "time"

type VerificationCode struct {
	CodeId       int64     `gorm:"column:code_id;primaryKey;autoIncrement"`
	UserId       int64     `gorm:"column:user_id"`
	Code         string    `gorm:"column:code"`
	Valid        bool      `gorm:"column:valid"`
	CreationTime time.Time `gorm:"column:creation_time"`
	ExpireTime   time.Time `gorm:"column:expire_time"`
}

func (VerificationCode) TableName() string {
	return "verification_codes"
}
