package entities

type Tag struct {
	TagId uint64 `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
}
