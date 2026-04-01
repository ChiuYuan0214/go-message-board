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
	var users []store.User
	err := r.db.Orm().Raw(`select u.user_id, u.username, ifnull(i.file_name, '') as user_image 
	from users u left join images i on i.user_id = u.user_id`).Scan(&users).Error
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]store.User, 0)
	}

	return users, nil
}
