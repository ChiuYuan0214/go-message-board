package types

import (
	"time"
)

type Article struct {
	ArticleId    uint64    `json:"articleId" gorm:"primaryKey"`
	UserId       uint64    `json:"userId" gorm:"column:user_id"`
	Author       string    `json:"author" gorm:"-"`
	AuthorImage  string    `json:"authorImage" gorm:"-"`
	Title        string    `json:"title" gorm:"column:title"`
	Content      string    `json:"content" gorm:"column:content"`
	TopCommentId uint64    `json:"topCommentId" gorm:"column:top_comment_id"`
	Edited       bool      `json:"edited" gorm:"column:edited"`
	ViewCount    uint32    `json:"viewCount" gorm:"column:view_count"`
	VoteUp       int32     `json:"voteUp" gorm:"-"`
	VoteDown     int32     `json:"voteDown" gorm:"-"`
	MyScore      int32     `json:"myScore" gorm:"-"`
	HasCollec    bool      `json:"hasCollec" gorm:"-"`
	PublishTime  time.Time `json:"publishTime" gorm:"column:publish_time"`
	CreationTime time.Time `json:"creationTime" gorm:"column:creation_time"`
	UpdateTime   time.Time `json:"updateTime" gorm:"column:update_time"`
	Tags         []string  `json:"tags" gorm:"-"`
}

type ArticleListData struct {
	ArticleId        uint64    `json:"articleId" gorm:"column:article_id"`
	UserId           uint64    `json:"userId" gorm:"column:user_id"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	Author           string    `json:"author" gorm:"column:username"`
	AuthorImage      string    `json:"authorImage" gorm:"column:file_name"`
	VoteUp           int32     `json:"voteUp"`
	VoteDown         int32     `json:"voteDown"`
	MyScore          int32     `json:"myScore"`
	HasCollec        bool      `json:"hasCollec"`
	CommentTitle     string    `json:"commentTitle" gorm:"column:title"`
	CommentContent   string    `json:"commentContent" gorm:"column:content"`
	CommentUser      string    `json:"commentUser" gorm:"column:username"`
	CommentUserImage string    `json:"commentUserImage" gorm:"column:file_name"`
	PublishTime      time.Time `json:"publishTime" gorm:"column:publish_time"`
	Tags             []string  `json:"tags"`
}
