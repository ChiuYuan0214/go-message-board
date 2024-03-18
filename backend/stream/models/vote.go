package models

type VoteRequest struct {
	Scores []uint8 `json:"scores"`
}

type Vote struct {
	SenderId uint64  `json:"senderId"`
	Scores   []uint8 `json:"scores"`
}

type SubjectRequest struct {
	Title     string   `json:"title"`
	Questions []string `json:"questions"`
}
