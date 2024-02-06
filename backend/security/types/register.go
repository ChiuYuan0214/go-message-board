package types

import "database/sql"

type RegisterData struct {
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Phone    sql.NullString `json:"phone"`
	Job      sql.NullString `json:"job"`
	Address  sql.NullString `json:"address"`
}
