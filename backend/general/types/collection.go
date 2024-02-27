package types

import "time"

type CollectionData struct {
	ArticleId   uint64    `json:"articleId"`
	UserId      uint64    `json:"userId"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	AuthorImage string    `json:"authorImage"`
	VoteUp      int32     `json:"voteUp"`
	VoteDown    int32     `json:"voteDown"`
	MyScore     int32     `json:"myScore"`
	HasCollec   bool      `json:"hasCollec"`
	PublishTime time.Time `json:"publishTime"`
}

type WriteCollectionData struct {
	ArticleId uint64 `json:"articleId"`
}
