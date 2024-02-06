package types

import "time"

type AddArticleData struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishTime time.Time `json:"publishTime"`
	Tags        []string  `json:"tags"`
}

type UpdateArticleData struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type AddCommentData struct {
	ArticleId int64  `json:"articleId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type UpdateCommentData struct {
	CommentId int64  `json:"commentId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
