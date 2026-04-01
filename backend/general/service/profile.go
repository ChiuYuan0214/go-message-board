package service

import (
	"general/repo"
	"general/types"
	"log"
	"net/http"
)

var _ Profile = (*ProfileImpl)(nil)

type ProfileImpl struct {
	profileRepo repo.Profile
}

func (s *ProfileImpl) Run() (err error) {
	return
}

func (s *ProfileImpl) Stop() {}

func (s *ProfileImpl) GetProfileWithId(userId uint64) (*types.Profile, int) {
	profile, err := s.profileRepo.GetByUserId(userId)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return &profile, http.StatusOK
}

func (s *ProfileImpl) GetProfileWithToken(userId uint64) (*types.SelfProfile, int) {
	profile, err := s.profileRepo.GetSelfById(userId)
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	return &profile, http.StatusOK
}
