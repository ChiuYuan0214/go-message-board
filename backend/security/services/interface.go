package services

import (
	"security/store"
	"security/types"
	"security/utils"
	"time"
)

type Auth interface {
	Login(email string, password string) (uint64, *types.Token)
	GenerateToken(userId uint64) *types.Token
	VerifyToken(userId uint64, token string) bool
	GetUserIdFromToken(srcToken string) uint64
}

type Register interface {
	CheckEmailExist(email string) bool
	AddNewUser(username string, email string, password string, phone string, job string, address string) int64
	InsertVerificationCode(userId int64, code int32, expireTime time.Time) int64
	InvalidateVerificationCodes(userId int64) bool
	InvalidateVerificationCodesByCodeId(codeId int64) bool
	GetActiveVerificationCode(userId uint64) (string, error)
	ScheduleCodeInvalidation(codeId int64, veriCode *utils.VerificationCode)
	ActivateUser(userId uint64) bool
	VerifyPasswordByEmail(email *string, password *string) int64
}

type Profile interface {
	VerifyPasswordByUserId(userId *uint64, password *string) bool
	UpdatePassword(userId *uint64, password *string) bool
	InsertProfileImageInfo(userId *uint64, fileName *string, desc *string) (string, int)
	UpdateColumnsById(data interface{}, id *uint64) (string, int)
}

type User interface {
	GetUsers(name string, userId uint64) []store.User
}
