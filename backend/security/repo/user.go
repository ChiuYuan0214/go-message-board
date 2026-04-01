package repo

import (
	"security/infra"
	"security/store"
)

var _ User = (*UserImpl)(nil)

type UserImpl struct {
	db infra.RDB
}

func NewUser(db infra.RDB) *UserImpl {
	return &UserImpl{
		db: db,
	}
}

func (r *UserImpl) ListUsers() ([]store.User, error) {
	rows, err := r.db.DB().Query(`select u.user_id, u.username, ifnull(i.file_name, '') from users u left join images i on i.user_id = u.user_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]store.User, 0)
	for rows.Next() {
		var user store.User
		if err = rows.Scan(&user.UserId, &user.UserName, &user.UserImage); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}
