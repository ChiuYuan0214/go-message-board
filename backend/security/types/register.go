package types

type RegisterData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Job      string `json:"job"`
	Address  string `json:"address"`
}
