package repo

import (
	"database/sql"
	"fmt"
	"security/infra"
	"security/store"
	"time"
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
	err = r.db.DB().QueryRow("select count(user_id) from users where email = ? ", email).Scan(&count)
	return count > 0, err
}

func (r *RegisterImpl) AddNewUser(username string, email string, password string, phone string, job string, address string) (int64, error) {
	sqlRes, err := r.db.DB().Exec(`insert into users (username, email, password, phone, job, address) 
	values (?, ?, ?, ?, ?, ?)`, username, email, password, phone, job, address)
	if err != nil {
		return 0, err
	}

	userId, err := sqlRes.LastInsertId()
	return userId, err
}

func (r *RegisterImpl) InsertVerificationCode(userId int64, code int32, expireTime time.Time) (int64, error) {
	sqlRes, err := r.db.DB().Exec(`insert into verification_codes (user_id, code, expire_time) 
	                              values (?, ?, ?)`, userId, fmt.Sprintf("%06d", code), expireTime)
	if err != nil {
		return 0, err
	}

	return sqlRes.LastInsertId()
}

func (r *RegisterImpl) InvalidateVerificationCodes(userId int64) error {
	_, err := r.db.DB().Exec(`update verification_codes set valid = false where user_id = ?`, userId)
	return err
}

func (r *RegisterImpl) InvalidateVerificationCodesByCodeId(codeId int64) error {
	_, err := r.db.DB().Exec(`update verification_codes set valid = false where code_id = ?`, codeId)
	return err
}

func (r *RegisterImpl) GetActiveVerificationCode(userId uint64) (code string, err error) {
	err = r.db.DB().QueryRow(`select code from verification_codes where user_id = ? and valid = true`, userId).Scan(&code)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return
}

func (r *RegisterImpl) GetUserById(userId uint64) (user store.User, err error) {
	err = r.db.DB().QueryRow(`select u.user_id, u.username, ifnull(i.file_name, '') from users u 
	left join images i on i.user_id = u.user_id where u.user_id = ?`, userId).Scan(&user.UserId, &user.UserName, &user.UserImage)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (r *RegisterImpl) ActivateUser(userId uint64) error {
	_, err := r.db.DB().Exec("update users set is_active = true where user_id = ?", userId)
	return err
}

func (r *RegisterImpl) GetCredentialStatusByEmail(email string) (userId int64, hashedPassword string, isActive bool, err error) {
	err = r.db.DB().QueryRow("select user_id, password, is_active from users where email = ? ", email).Scan(&userId, &hashedPassword, &isActive)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
