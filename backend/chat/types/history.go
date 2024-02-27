package types

type History struct {
	Event         string    `json:"event"`
	TargetUserId  uint64    `json:"targetUserId"`
	UserHistory   []Message `json:"userHistory"`
	TargetHistory []Message `json:"targetHistory"`
}
