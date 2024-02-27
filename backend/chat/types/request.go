package types

type RequestEvent struct {
	UserId       uint64 `json:"userId"`
	TargetUserId uint64 `json:"targetUserId"`
	Type         string `json:"type"`
	Content      string `json:"content"`
}
