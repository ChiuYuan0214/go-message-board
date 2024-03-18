package types

type AddArticleData struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	PublishTime string   `json:"publishTime"`
	Tags        []string `json:"tags"`
}

type UpdateArticleData struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type AddCommentData struct {
	ArticleId uint64 `json:"articleId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type UpdateCommentData struct {
	CommentId uint64 `json:"commentId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type FollowData struct {
	Followee uint64 `json:"followee"`
}

type FollowerData struct {
	Follower uint64 `json:"follower"`
}
