package entities

type Follower struct {
	UserId     uint64 `json:"userId" gorm:"primaryKey"`
	FollowerId uint64 `json:"followerId" gorm:"primaryKey"`
}
