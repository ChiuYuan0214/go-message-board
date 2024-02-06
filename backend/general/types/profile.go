package types

type Profile struct {
	UserId   uint64 `json:"userId"`
	Username string `json:"username"`
	Job      string `json:"job"`
	IsActive bool   `json:"isActive"`
}

type SelfProfile struct {
	UserId       uint64 `json:"userId"`
	Username     string `json:"username"`
	IsActive     bool   `json:"isActive"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Job          string `json:"job"`
	Address      string `json:"address"`
	CreationTime string `json:"creationTime"`
}
