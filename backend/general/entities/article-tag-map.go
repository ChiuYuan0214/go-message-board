package entities

type ArticleTagMap struct {
	ArticleId uint64 `gorm:"primaryKey"`
	TagId     uint64 `gorm:"primaryKey"`
}
