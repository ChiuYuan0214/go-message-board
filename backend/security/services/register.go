package services

import (
	"database/sql"
	"fmt"
	"security/utils"
	"time"
)

// type ErrorHandler struct {
// 	result interface{}
// }

// func (eh *ErrorHandler) Check(checker func(any) bool) {
// 	if !checker(eh.result) {
// 	}
// }

func CheckEmailExist(email string) bool {
	row := connPool.QueryRow("select count(user_id) from users where email = ? ", email)
	var count int64
	err := row.Scan(&count)
	if err != nil || count > 0 {
		return true
	}
	return false
}

func AddNewUser(username string, email string, password string,
	phone sql.NullString, job sql.NullString, address sql.NullString) int64 {
	sqlRes, err := connPool.Exec(`insert into users (username, email, password, phone, job, address) 
	values (?, ?, ?, ?, ?, ?)`, username, email, password, phone, job, address)

	count, err := sqlRes.RowsAffected()
	if err != nil || count < 1 {
		return 0
	}
	userId, _ := sqlRes.LastInsertId()
	return userId
}

func InsertVerificationCode(userId int64, code int32, expireTime time.Time) int64 {
	sqlRes, err := connPool.Exec(`insert into verification_codes (user_id, code, expire_time) 
	                              values (?, ?, ?)`, userId, fmt.Sprintf("%06d", code), expireTime)
	id, err := sqlRes.LastInsertId()
	if err != nil {
		return 0
	}
	return id
}

func InvalidateVerificationCodes(userId int64) bool {
	sqlRes, err := connPool.Exec(`update verification_codes set valid = false where user_id = ?`, userId)
	_, err = sqlRes.RowsAffected()
	if err != nil {
		return false
	}
	return true
}

func InvalidateVerificationCodesByCodeId(codeId int64) bool {
	sqlRes, err := connPool.Exec(`update verification_codes set valid = false where code_id = ?`, codeId)
	_, err = sqlRes.RowsAffected()
	if err != nil {
		return false
	}
	return true
}

func ScheduleCodeInvalidation(codeId int64, veriCode *utils.VerificationCode) {
	go func() {
		time.Sleep(time.Second)
		currentTime := time.Now().Unix()
		if currentTime >= (*veriCode).ExpireTime.Unix() {
			InvalidateVerificationCodesByCodeId(codeId)
			return
		}
	}()
}

func ActivateUser(userId int64) bool {
	_, err := connPool.Exec("update users set is_active = true where user_id = ?", userId)
	if err != nil {
		return false
	}
	return true
}

func VerifyPasswordByEmail(email *string, password *string) int64 {
	var userId int64
	var hashedPassword string
	var isActive bool
	row := connPool.QueryRow("select user_id, password, is_active from users where email = ? ", *email)
	err := row.Scan(&userId, &hashedPassword, &isActive)
	if err != nil || !utils.VerifyPassword(&hashedPassword, password) {
		return 0
	}
	if isActive {
		return -1
	}

	return userId
}
