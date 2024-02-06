package types

type CommentListData struct {
	CommentId    int64  `json:"commentId"`
	UserId       int64  `json:"userId"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreationTime string `json:"updateTime"`
	VoteUp       int16  `json:"voteUp"`
	VoteDown     int16  `json:"voteDown"`
}
