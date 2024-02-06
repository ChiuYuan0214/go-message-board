package services

import (
	"database/sql"
	"general/types"
	"net/http"
	"strconv"
)

func GetProfileWithId(userId string) (*types.Profile, int) {
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || id == 0 {
		return nil, http.StatusBadRequest
	}
	var profile types.Profile
	var job sql.NullString
	var isActive sql.NullBool
	row := connPool.QueryRow("select user_id, username, job, is_active from users where user_id = ?", userId)
	err = row.Scan(&profile.UserId, &profile.Username, &job, &isActive)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	if job.Valid {
		val, _ := job.Value()
		profile.Job = val.(string)
	}
	if isActive.Valid {
		val, _ := isActive.Value()
		profile.IsActive = val.(bool)
	}

	return &profile, http.StatusOK
}

func GetProfileWithToken(userId int64) (*types.SelfProfile, int) {
	var profile types.SelfProfile
	var isActive sql.NullBool
	var phone sql.NullString
	var job sql.NullString
	var address sql.NullString

	row := connPool.QueryRow(
		"select user_id, username, is_active, email, phone, job, address, creation_time from users where user_id = ?", userId)
	err := row.Scan(&profile.UserId, &profile.Username, &isActive, &profile.Email, &phone, &job, &address, &profile.CreationTime)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	if isActive.Valid {
		val, _ := isActive.Value()
		profile.IsActive = val.(bool)
	}
	if phone.Valid {
		val, _ := phone.Value()
		profile.Phone = val.(string)
	}
	if job.Valid {
		val, _ := job.Value()
		profile.Job = val.(string)
	}
	if address.Valid {
		val, _ := address.Value()
		profile.Address = val.(string)
	}

	return &profile, http.StatusOK
}
