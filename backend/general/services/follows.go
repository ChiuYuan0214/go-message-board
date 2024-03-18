package services

import (
	"general/types"
	"log"
)

func GetFollows(userId uint64) []types.Follower {
	data := []types.Follower{}
	err := db.Table("followers f").Select("f.user_id, u.username, i.file_name").
		Joins("inner join users u on u.user_id = f.user_id").
		Joins("left join images i on i.user_id = u.user_id").
		Where("f.follower_id = ?", userId).Find(&data).Error
	if err != nil {
		log.Println(err)
		return []types.Follower{}
	}

	return data
}
