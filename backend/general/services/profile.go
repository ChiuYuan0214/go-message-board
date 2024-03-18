package services

import (
	"general/types"
	"log"
	"net/http"
)

func GetProfileWithId(userId uint64) (*types.Profile, int) {
	var profile types.Profile
	err := db.Raw(`select u.user_id, u.username, u.job, u.is_active, i.file_name as imagePath, 
	ifnull((select count(article_id) from articles a where a.user_id = u.user_id), 0) as articleCount, 
	ifnull((select count(comment_id) from comments c where c.user_id = u.user_id), 0) as commentCount, 
	ifnull((select count(vote_id) from votes v inner join articles a 
	on a.article_id = v.source_id and vote_type = 'article' and score = 1 and a.user_id = u.user_id) + 
	(select count(vote_id) from votes v inner join comments c 
	on c.comment_id = v.source_id and vote_type = 'comment' and score = 1 and c.user_id = u.user_id), 0) as upVoteCount 
	from users u left join images i on i.user_id = u.user_id where u.user_id = ?`, userId).Scan(&profile).Error
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return &profile, http.StatusOK
}

func GetProfileWithToken(userId uint64) (*types.SelfProfile, int) {
	var profile types.SelfProfile

	err := db.Raw(`select u.user_id, u.username, u.is_active, u.email, u.phone, u.job, u.address, i.file_name, u.creation_time, 
		ifnull((select count(article_id) from articles a where a.user_id = u.user_id), 0) as articleCount, 
	    ifnull((select count(comment_id) from comments c where c.user_id = u.user_id), 0) as commentCount, 
	    ifnull((select count(vote_id) from votes v inner join articles a 
	    on a.article_id = v.source_id and vote_type = 'article' and score = 1 and a.user_id = u.user_id) + 
	    (select count(vote_id) from votes v inner join comments c 
	    on c.comment_id = v.source_id and vote_type = 'comment' and score = 1 and c.user_id = u.user_id), 0) as upVoteCount 
		from users u left join images i on i.user_id = u.user_id where u.user_id = ?`, userId).Scan(&profile).Error
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	return &profile, http.StatusOK
}
