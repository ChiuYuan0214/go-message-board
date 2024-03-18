package models

type FeedbackRequest struct {
	Content string `json:"content"`
}

type Feedback struct {
	SenderId uint64 `json:"senderId"`
	Content  string `json:"content"`
}

type ResponseFeedbackRequest struct {
	ReceiverId uint64 `json:"receiverId"`
	Content    string `json:"content"`
}
