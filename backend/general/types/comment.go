package types

type CommentListData struct {
	CommentId      int64  `json:"commentId" gorm:"column:comment_id"`
	UserId         int64  `json:"userId" gorm:"column:user_id"`
	Commenter      string `json:"commenter"`
	CommenterImage string `json:"commenterImage"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	CreationTime   string `json:"updateTime"`
	VoteUp         int16  `json:"voteUp"`
	VoteDown       int16  `json:"voteDown"`
}
