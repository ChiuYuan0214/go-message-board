package types

import "time"

type User struct {
	UserId       uint64    `json:"userId"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreationTime time.Time `json:"creationTime"`
}
