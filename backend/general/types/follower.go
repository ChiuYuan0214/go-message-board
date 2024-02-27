package types

type Follower struct {
	UserId    int64  `json:"userId"`
	Username  string `json:"username"`
	UserImage string `json:"userImage"`
}
