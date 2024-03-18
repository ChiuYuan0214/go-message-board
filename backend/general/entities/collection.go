package entities

type Collection struct {
	UserId    uint64 `gorm:"primaryKey"`
	ArticleId uint64 `gorm:"primaryKey"`
}
