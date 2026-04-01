package services

import (
	"security/repo"
	"security/store"
	"security/utils"
	"time"
)

var _ Register = (*RegisterImpl)(nil)

type RegisterImpl struct {
	registerRepo repo.Register
	usersStore   *store.UsersStore
}

func NewRegister(registerRepo repo.Register, usersStore *store.UsersStore) *RegisterImpl {
	return &RegisterImpl{
		registerRepo: registerRepo,
		usersStore:   usersStore,
	}
}

func (s *RegisterImpl) CheckEmailExist(email string) bool {
	exists, err := s.registerRepo.CheckEmailExist(email)
	return err != nil || exists
}

func (s *RegisterImpl) AddNewUser(username string, email string, password string,
	phone string, job string, address string) int64 {
	userId, err := s.registerRepo.AddNewUser(username, email, password, phone, job, address)
	if err != nil || userId < 1 {
		return 0
	}

	return userId
}

func (s *RegisterImpl) InsertVerificationCode(userId int64, code int32, expireTime time.Time) int64 {
	id, err := s.registerRepo.InsertVerificationCode(userId, code, expireTime)
	if err != nil {
		return 0
	}
	return id
}

func (s *RegisterImpl) InvalidateVerificationCodes(userId int64) bool {
	return s.registerRepo.InvalidateVerificationCodes(userId) == nil
}

func (s *RegisterImpl) InvalidateVerificationCodesByCodeId(codeId int64) bool {
	return s.registerRepo.InvalidateVerificationCodesByCodeId(codeId) == nil
}

func (s *RegisterImpl) GetActiveVerificationCode(userId uint64) (string, error) {
	return s.registerRepo.GetActiveVerificationCode(userId)
}

func (s *RegisterImpl) ScheduleCodeInvalidation(codeId int64, veriCode *utils.VerificationCode) {
	go func() {
		time.Sleep(time.Second)
		currentTime := time.Now().Unix()
		if currentTime >= (*veriCode).ExpireTime.Unix() {
			s.InvalidateVerificationCodesByCodeId(codeId)
			return
		}
	}()
}

func (s *RegisterImpl) ActivateUser(userId uint64) bool {
	user, err := s.registerRepo.GetUserById(userId)
	if err != nil {
		return false
	}
	if user.UserId != 0 {
		s.usersStore.AddUser(user)
	}

	return s.registerRepo.ActivateUser(userId) == nil
}

func (s *RegisterImpl) VerifyPasswordByEmail(email *string, password *string) int64 {
	userId, hashedPassword, isActive, err := s.registerRepo.GetCredentialStatusByEmail(*email)
	if err != nil || !utils.VerifyPassword(&hashedPassword, password) {
		return 0
	}
	if isActive {
		return -1
	}

	return userId
}
