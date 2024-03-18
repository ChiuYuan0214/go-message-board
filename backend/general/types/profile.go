package types

type Profile struct {
	UserId       uint64 `json:"userId" gorm:"column:user_id"`
	Username     string `json:"username"`
	Job          string `json:"job"`
	IsActive     bool   `json:"isActive" gorm:"column:is_active"`
	ImagePath    string `json:"imagePath"`
	ArticleCount int32  `json:"articleCount"`
	CommentCount int32  `json:"commentCount"`
	UpVoteCount  int32  `json:"upVoteCount"`
}

type SelfProfile struct {
	UserId       uint64 `json:"userId"`
	Username     string `json:"username"`
	IsActive     bool   `json:"isActive"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Job          string `json:"job"`
	Address      string `json:"address"`
	ImagePath    string `json:"imagePath"`
	CreationTime string `json:"creationTime"`
	ArticleCount int32  `json:"articleCount"`
	CommentCount int32  `json:"commentCount"`
	UpVoteCount  int32  `json:"upVoteCount"`
}
