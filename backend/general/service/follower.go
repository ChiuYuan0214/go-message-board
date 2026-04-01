package service

import (
	"general/repo"
	"general/types"
	"log"
)

var _ Follower = (*FollowerImpl)(nil)

type FollowerImpl struct {
	followerRepo repo.Follower
}

func (s *FollowerImpl) Run() (err error) {
	return
}

func (s *FollowerImpl) Stop() {}

func (s *FollowerImpl) RemoveFollower(userId uint64, follower uint64) bool {
	return s.followerRepo.DeleteByUserIdAndFollowerId(userId, follower) == nil
}

func (s *FollowerImpl) GetFollowers(userId uint64) (data []types.Follower) {
	data, err := s.followerRepo.GetFollowersByUserId(userId)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
