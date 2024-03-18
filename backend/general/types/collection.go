package types

import "time"

type CollectionData struct {
	ArticleId   uint64    `json:"articleId" gorm:"column:article_id"`
	UserId      uint64    `json:"userId" gorm:"column:user_id"`
	Title       string    `json:"title" gorm:"column:title"`
	Content     string    `json:"content" gorm:"column:content"`
	Author      string    `json:"author" gorm:"column:author"`
	AuthorImage string    `json:"authorImage"  gorm:"column:authorImage"`
	VoteUp      int32     `json:"voteUp" gorm:"column:voteUp"`
	VoteDown    int32     `json:"voteDown" gorm:"column:voteDown"`
	MyScore     int32     `json:"myScore" gorm:"column:myScore"`
	HasCollec   bool      `json:"hasCollec" gorm:"column:hasCollec"`
	PublishTime time.Time `json:"publishTime" gorm:"column:publishTime"`
}

type WriteCollectionData struct {
	ArticleId uint64 `json:"articleId"`
}
