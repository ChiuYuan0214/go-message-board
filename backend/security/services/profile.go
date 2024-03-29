package services

import (
	"net/http"
	"security/utils"
)

func VerifyPasswordByUserId(userId *uint64, password *string) bool {
	var hashedPassword string
	row := connPool.QueryRow("select password from users where user_id = ? ", *userId)
	err := row.Scan(&hashedPassword)
	if err != nil || !utils.VerifyPassword(&hashedPassword, password) {
		return false
	}

	return true
}

func UpdatePassword(userId *uint64, password *string) bool {
	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return false
	}

	sqlRes, err := connPool.Exec("update users set password = ? where user_id = ?", hashedPassword, *userId)
	if err != nil {
		return false
	}
	count, _ := sqlRes.RowsAffected()

	return count == 1
}

func InsertProfileImageInfo(userId *uint64, fileName *string, desc *string) (string, int) {
	var count int64
	var err error
	row := connPool.QueryRow("select count(user_id) from images where user_id = ?", *userId)
	if err = row.Scan(&count); err != nil {
		return "failed to insert image info.", http.StatusInternalServerError
	}

	if count > 0 {
		_, err = connPool.Exec("update images set file_name = ?, descript = ? where user_id = ?", *fileName, *desc, *userId)
	} else {
		_, err = connPool.Exec("insert into images (user_id, file_name, descript) values (?, ?, ?)", *userId, *fileName, *desc)
	}
	if err != nil {
		return "failed to insert image info.", http.StatusInternalServerError
	}

	return "", 0
}
