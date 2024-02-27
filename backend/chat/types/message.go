package types

type Message struct {
	Event        string `json:"event"`
	UserId       uint64 `json:"userId"`
	TargetUserId uint64 `json:"targetUserId"`
	Content      string `json:"content"`
	Time         int64  `json:"time"`
	HasSync      bool   `json:"-"`
	Ref          uint8  `json:"-"`
}

type ServerMessage struct {
	Event   string `json:"event"`
	Content string `json:"content"`
}
