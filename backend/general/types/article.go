package types

import (
	"time"
)

type Article struct {
	ArticleId    uint64    `json:"articleId"`
	UserId       uint64    `json:"userId"`
	Author       string    `json:"author"`
	AuthorImage  string    `json:"authorImage"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	TopCommentId uint64    `json:"topCommentId"`
	Edited       bool      `json:"edited"`
	ViewCount    uint32    `json:"viewCount"`
	VoteUp       int32     `json:"voteUp"`
	VoteDown     int32     `json:"voteDown"`
	MyScore      int32     `json:"myScore"`
	HasCollec    bool      `json:"hasCollec"`
	PublishTime  time.Time `json:"publishTime"`
	CreationTime time.Time `json:"creationTime"`
	UpdateTime   time.Time `json:"updateTime"`
	Tags         []string  `json:"tags"`
}

type ArticleListData struct {
	ArticleId        uint64    `json:"articleId"`
	UserId           uint64    `json:"userId"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	Author           string    `json:"author"`
	AuthorImage      string    `json:"authorImage"`
	VoteUp           int32     `json:"voteUp"`
	VoteDown         int32     `json:"voteDown"`
	MyScore          int32     `json:"myScore"`
	HasCollec        bool      `json:"hasCollec"`
	CommentTitle     string    `json:"commentTitle"`
	CommentContent   string    `json:"commentContent"`
	CommentUser      string    `json:"commentUser"`
	CommentUserImage string    `json:"commentUserImage"`
	PublishTime      time.Time `json:"publishTime"`
	Tags             []string  `json:"tags"`
}
