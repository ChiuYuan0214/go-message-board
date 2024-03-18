package services

import (
	"general/entities"
	"log"
)

func AddFollow(userId uint64, followee uint64) bool {
	var count int64
	follower := entities.Follower{}
	err := db.Model(&follower).Where("user_id = ? and follower_id = ?", followee, userId).Count(&count).Error
	if err != nil {
		log.Println(err)
		return false
	}
	if count > 0 {
		return true
	}
	follower.UserId = followee
	follower.FollowerId = userId
	err = db.Create(&follower).Error
	if err != nil {
		log.Println(err)
	}
	return err == nil
}

func RemoveFollow(userId uint64, followee uint64) bool {
	var count int64
	err := db.Model(&entities.Follower{}).Where("user_id = ? and follower_id = ?", followee, userId).Count(&count).Error
	if err != nil {
		log.Println(err)
		return false
	}
	if count == 0 {
		return true
	}
	err = db.Where("user_id = ? and follower_id = ?", followee, userId).Delete(&entities.Follower{}).Error
	if err != nil {
		log.Println(err)
	}
	return err == nil
}
