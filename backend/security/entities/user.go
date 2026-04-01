package entities

import "time"

type User struct {
	UserId       uint64    `gorm:"column:user_id;primaryKey;autoIncrement"`
	UserName     string    `gorm:"column:username"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	Phone        string    `gorm:"column:phone"`
	Job          string    `gorm:"column:job"`
	Address      string    `gorm:"column:address"`
	IsActive     bool      `gorm:"column:is_active"`
	CreationTime time.Time `gorm:"column:creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}

func (User) TableName() string {
	return "users"
}
