package repo

import (
	"fmt"
	"security/entities"
	"security/infra"
	"security/store"
	"time"

	"gorm.io/gorm"
)

var _ Register = (*RegisterImpl)(nil)

type RegisterImpl struct {
	db infra.RDB
}

func NewRegister(db infra.RDB) *RegisterImpl {
	return &RegisterImpl{
		db: db,
	}
}

func (r *RegisterImpl) CheckEmailExist(email string) (exists bool, err error) {
	var count int64
	err = r.db.Orm().Model(new(entities.User)).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *RegisterImpl) AddNewUser(username string, email string, password string, phone string, job string, address string) (int64, error) {
	user := entities.User{
		UserName: username,
		Email:    email,
		Password: password,
		Phone:    phone,
		Job:      job,
		Address:  address,
	}

	err := r.db.Orm().Create(&user).Error
	return int64(user.UserId), err
}

func (r *RegisterImpl) InsertVerificationCode(userId int64, code int32, expireTime time.Time) (int64, error) {
	record := entities.VerificationCode{
		UserId:     userId,
		Code:       fmt.Sprintf("%06d", code),
		ExpireTime: expireTime,
	}

	err := r.db.Orm().Create(&record).Error
	return record.CodeId, err
}

func (r *RegisterImpl) InvalidateVerificationCodes(userId int64) error {
	return r.db.Orm().
		Model(new(entities.VerificationCode)).
		Where("user_id = ?", userId).
		Update("valid", false).Error
}

func (r *RegisterImpl) InvalidateVerificationCodesByCodeId(codeId int64) error {
	return r.db.Orm().
		Model(new(entities.VerificationCode)).
		Where("code_id = ?", codeId).
		Update("valid", false).Error
}

func (r *RegisterImpl) GetActiveVerificationCode(userId uint64) (code string, err error) {
	var record entities.VerificationCode
	err = r.db.Orm().
		Select("code").
		Where("user_id = ? and valid = true", userId).
		First(&record).Error
	if err == gorm.ErrRecordNotFound {
		return "", nil
	}

	code = record.Code
	return
}

func (r *RegisterImpl) GetUserById(userId uint64) (user store.User, err error) {
	err = r.db.Orm().Raw(`select u.user_id, u.username, ifnull(i.file_name, '') as user_image from users u 
	left join images i on i.user_id = u.user_id where u.user_id = ?`, userId).Scan(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (r *RegisterImpl) ActivateUser(userId uint64) error {
	return r.db.Orm().
		Model(new(entities.User)).
		Where("user_id = ?", userId).
		Update("is_active", true).Error
}

func (r *RegisterImpl) GetCredentialStatusByEmail(email string) (userId int64, hashedPassword string, isActive bool, err error) {
	var user entities.User
	err = r.db.Orm().
		Select("user_id", "password", "is_active").
		Where("email = ?", email).
		First(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}

	userId = int64(user.UserId)
	hashedPassword = user.Password
	isActive = user.IsActive
	return
}
