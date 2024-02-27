package services

import (
	"database/sql"
	"general/types"
	"log"
)

func GetFollows(userId int64) *[]types.Follower {
	data := &[]types.Follower{}
	rows, err := connPool.Query(`select f.user_id, u.username, i.file_name from followers f 
	inner join users u on u.user_id = f.user_id left join images i on i.user_id = u.user_id where f.follower_id = ?`, userId)
	if err != nil {
		log.Println(err)
		return data
	}
	for rows.Next() {
		var follower types.Follower
		var userImage sql.NullString
		err := rows.Scan(&follower.UserId, &follower.Username, &userImage)
		if err != nil {
			log.Println(err)
			continue
		}
		if userImage.Valid {
			val, _ := userImage.Value()
			follower.UserImage = val.(string)
		}
		*data = append(*data, follower)
	}
	return data
}
