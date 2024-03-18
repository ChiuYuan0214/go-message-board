package services

import (
	"general/entities"
	"general/types"
	"log"
)

func RemoveFollower(userId uint64, follower uint64) bool {
	result := db.Where("user_id = ? and follower_id = ?", userId, follower).Delete(&entities.Follower{})
	return result.Error != nil && result.RowsAffected > 0
}

func GetFollowers(userId uint64) []types.Follower {
	data := []types.Follower{}
	err := db.Table("followers f").
		Select("f.follower_id, u.username, i.file_name").
		Joins("inner join users u on u.user_id = f.follower_id").
		Joins("left join images i on i.user_id = u.user_id").
		Where("f.user_id = ?", userId).Find(&data).Error
	if err != nil {
		log.Println(err)
		return []types.Follower{}
	}
	return data
}
