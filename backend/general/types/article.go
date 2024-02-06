package types

import (
	"time"
)

type Article struct {
	ArticleId    uint64    `json:"articleId"`
	UserId       uint64    `json:"userId"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	TopCommentId uint64    `json:"topCommentId"`
	Edited       bool      `json:"edited"`
	ViewCount    uint32    `json:"viewCount"`
	VoteUp       int32     `json:"voteUp"`
	VoteDown     int32     `json:"voteDown"`
	PublishTime  time.Time `json:"publishTime"`
	CreationTime time.Time `json:"creationTime"`
	UpdateTime   time.Time `json:"updateTime"`
	Tags         []string  `json:"tags"`
}

type ArticleListData struct {
	ArticleId    uint64    `json:"articleId"`
	UserId       uint64    `json:"userId"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	TopCommentId uint64    `json:"topCommentId"`
	VoteUp       int32     `json:"voteUp"`
	VoteDown     int32     `json:"voteDown"`
	UpdateTime   time.Time `json:"updateTime"`
	Tags         []string  `json:"tags"`
}
