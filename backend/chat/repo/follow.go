package repo

import (
	"chat/infra"
	"database/sql"
)

var _ Follow = (*FollowImpl)(nil)

type FollowImpl struct {
	db infra.RDB
}

func (r *FollowImpl) Run() (err error) {
	return
}

func (r *FollowImpl) Stop() {}

func (r *FollowImpl) GetFollowerIDs(userId uint64) (list []uint64, err error) {
	var rows *sql.Rows
	rows, err = r.db.DB().Query(`select follower_id from followers where user_id = ?`, userId)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var followerID uint64
		err = rows.Scan(&followerID)
		if err != nil {
			return
		}
		list = append(list, followerID)
	}

	err = rows.Err()

	return
}

func (r *FollowImpl) GetFollowIDs(userId uint64) (list []uint64, err error) {
	var rows *sql.Rows
	rows, err = r.db.DB().Query(`select user_id from followers where follower_id = ?`, userId)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var followID uint64
		err = rows.Scan(&followID)
		if err != nil {
			return
		}
		list = append(list, followID)
	}

	err = rows.Err()

	return
}
