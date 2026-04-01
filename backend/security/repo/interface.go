package repo

import (
	"security/store"
	"security/types"
	"time"
)

type Auth interface {
	GetLoginCredentialByEmail(email string) (uint64, string, error)
	SetToken(userId uint64, token types.Token) error
	GetToken(userId uint64) (string, error)
}

type Register interface {
	CheckEmailExist(email string) (bool, error)
	AddNewUser(username string, email string, password string, phone string, job string, address string) (int64, error)
	InsertVerificationCode(userId int64, code int32, expireTime time.Time) (int64, error)
	InvalidateVerificationCodes(userId int64) error
	InvalidateVerificationCodesByCodeId(codeId int64) error
	GetActiveVerificationCode(userId uint64) (string, error)
	GetUserById(userId uint64) (store.User, error)
	ActivateUser(userId uint64) error
	GetCredentialStatusByEmail(email string) (int64, string, bool, error)
}

type Profile interface {
	GetPasswordByUserId(userId uint64) (string, error)
	UpdatePassword(userId uint64, password string) (int64, error)
	UpsertProfileImageInfo(userId uint64, fileName string, desc string) error
	UpdateColumnsById(data interface{}, id uint64) (int64, error)
}

type User interface {
	ListUsers() ([]store.User, error)
}
