package services

import "log"

func AddFollow(userId int64, followee uint64) bool {
	var count uint8
	row := connPool.QueryRow(`select count(user_id) from followers where user_id = ? and follower_id = ?`, followee, userId)
	row.Scan(&count)
	if count > 0 {
		return true
	}
	sqlRes, err := connPool.Exec(`insert into followers (user_id, follower_id) values (?, ?)`, followee, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	_, err = sqlRes.RowsAffected()
	return err == nil
}

func RemoveFollow(userId int64, followee uint64) bool {
	var count uint8
	row := connPool.QueryRow(`select count(user_id) from followers where user_id = ? and follower_id = ?`, followee, userId)
	row.Scan(&count)
	if count == 0 {
		return true
	}
	sqlRes, err := connPool.Exec(`delete from followers where user_id = ? and follower_id = ?`, followee, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	_, err = sqlRes.RowsAffected()
	return err == nil
}
