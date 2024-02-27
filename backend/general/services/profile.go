package services

import (
	"database/sql"
	"general/types"
	"log"
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
	var imagePath sql.NullString
	row := connPool.QueryRow(`select u.user_id, u.username, u.job, u.is_active, i.file_name, 
	ifnull((select count(article_id) from articles a where a.user_id = u.user_id), 0) as articleCount, 
	ifnull((select count(comment_id) from comments c where c.user_id = u.user_id), 0) as commentCount, 
	ifnull((select count(vote_id) from votes v inner join articles a 
	on a.article_id = v.source_id and vote_type = 'article' and score = 1 and a.user_id = u.user_id) + 
	(select count(vote_id) from votes v inner join comments c 
	on c.comment_id = v.source_id and vote_type = 'comment' and score = 1 and c.user_id = u.user_id), 0) as upVoteCount 
	from users u left join images i on i.user_id = u.user_id where u.user_id = ?`, userId)
	err = row.Scan(&profile.UserId, &profile.Username, &job, &isActive, &imagePath,
		&profile.ArticleCount, &profile.CommentCount, &profile.UpVoteCount)
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
	if imagePath.Valid {
		val, _ := imagePath.Value()
		profile.ImagePath = val.(string)
	}

	return &profile, http.StatusOK
}

func GetProfileWithToken(userId int64) (*types.SelfProfile, int) {
	var profile types.SelfProfile
	var isActive sql.NullBool
	var phone sql.NullString
	var job sql.NullString
	var address sql.NullString
	var imagePath sql.NullString

	row := connPool.QueryRow(
		`select u.user_id, u.username, u.is_active, u.email, u.phone, u.job, u.address, i.file_name, u.creation_time, 
		ifnull((select count(article_id) from articles a where a.user_id = u.user_id), 0) as articleCount, 
	    ifnull((select count(comment_id) from comments c where c.user_id = u.user_id), 0) as commentCount, 
	    ifnull((select count(vote_id) from votes v inner join articles a 
	    on a.article_id = v.source_id and vote_type = 'article' and score = 1 and a.user_id = u.user_id) + 
	    (select count(vote_id) from votes v inner join comments c 
	    on c.comment_id = v.source_id and vote_type = 'comment' and score = 1 and c.user_id = u.user_id), 0) as upVoteCount 
		from users u left join images i on i.user_id = u.user_id where u.user_id = ?`, userId)
	err := row.Scan(&profile.UserId, &profile.Username, &isActive, &profile.Email,
		&phone, &job, &address, &imagePath, &profile.CreationTime,
		&profile.ArticleCount, &profile.CommentCount, &profile.UpVoteCount)
	if err != nil {
		log.Println(err)
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
	if imagePath.Valid {
		val, _ := imagePath.Value()
		profile.ImagePath = val.(string)
	}

	return &profile, http.StatusOK
}
