package types

type Notification struct {
	Event  string `json:"event"`
	UserId uint64 `json:"userId"`
}

type UserInfoList struct {
	Event string   `json:"event"`
	List  []uint64 `json:"list"`
}
