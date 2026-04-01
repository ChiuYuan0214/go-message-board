package services

import (
	"net/http"
	"security/repo"
	"security/utils"
)

var _ Profile = (*ProfileImpl)(nil)

type ProfileImpl struct {
	profileRepo repo.Profile
}

func NewProfile(profileRepo repo.Profile) *ProfileImpl {
	return &ProfileImpl{
		profileRepo: profileRepo,
	}
}

func (s *ProfileImpl) VerifyPasswordByUserId(userId *uint64, password *string) bool {
	hashedPassword, err := s.profileRepo.GetPasswordByUserId(*userId)
	if err != nil || !utils.VerifyPassword(&hashedPassword, password) {
		return false
	}

	return true
}

func (s *ProfileImpl) UpdatePassword(userId *uint64, password *string) bool {
	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return false
	}

	count, err := s.profileRepo.UpdatePassword(*userId, hashedPassword)
	if err != nil {
		return false
	}

	return count == 1
}

func (s *ProfileImpl) InsertProfileImageInfo(userId *uint64, fileName *string, desc *string) (string, int) {
	if err := s.profileRepo.UpsertProfileImageInfo(*userId, *fileName, *desc); err != nil {
		return "failed to insert image info.", http.StatusInternalServerError
	}

	return "", 0
}
