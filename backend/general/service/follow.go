package service

import (
	"errors"
	"general/repo"
	"log"

	"gorm.io/gorm"
)

var _ Follow = (*FollowImpl)(nil)

type FollowImpl struct {
	followerRepo repo.Follower
}

func (s *FollowImpl) Run() (err error) {
	return
}

func (s *FollowImpl) Stop() {}

func (s *FollowImpl) AddFollow(userId uint64, followee uint64) bool {
	_, err := s.followerRepo.GetByUserIdAndFollowerId(followee, userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		return false
	}
	if err == nil {
		return true
	}

	s.followerRepo.Create(followee, userId)
	if err != nil {
		log.Println(err)
	}

	return err == nil
}

func (s *FollowImpl) RemoveFollow(userId uint64, followee uint64) bool {
	_, err := s.followerRepo.GetByUserIdAndFollowerId(followee, userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		return false
	}
	if err == nil {
		return true
	}

	err = s.followerRepo.DeleteByUserIdAndFollowerId(followee, userId)
	if err != nil {
		log.Println(err)
	}
	return err == nil
}
