package services

import (
	"database/sql"
	"log"
	"net/http"
)

func Vote(userId int64, sourceId int64, score int16, voteType *string) (string, int64) {
	var row *sql.Row
	if *voteType == "article" {
		row = connPool.QueryRow("select count(*) from articles where article_id = ?", sourceId)
	} else {
		row = connPool.QueryRow("select count(*) from comments where comment_id = ?", sourceId)
	}

	var count int
	var voteId int64
	if err := row.Scan(&count); err != nil || count == 0 {
		return "source not exist", http.StatusBadRequest
	}
	row = connPool.QueryRow("select vote_id from votes where user_id = ? and source_id = ?", userId, sourceId)
	if err := row.Scan(&voteId); err == nil {
		if UpdateVote(userId, voteId, score) {
			return "", voteId
		} else {
			return "failed to execute query", http.StatusInternalServerError
		}
	}

	sqlRes, err := connPool.Exec("insert into votes (user_id, source_id, score, vote_type) values (?, ?, ?, ?)", userId, sourceId, score, *voteType)
	if err != nil {
		log.Println(err)
		return "failed to execute query", http.StatusInternalServerError
	}

	id, _ := sqlRes.LastInsertId()
	return "", id
}

func UpdateVote(userId int64, voteId int64, score int16) bool {
	row := connPool.QueryRow("select user_id, score from votes where vote_id = ?", voteId)
	var actualUserId int64
	var prevScore int16
	err := row.Scan(&actualUserId, &prevScore)
	if err != nil {
		log.Println(err)
		return false
	}
	if userId != actualUserId {
		return false
	}
	if prevScore == score {
		score = 0
	}

	_, err = connPool.Exec("update votes set score = ? where vote_id = ?", score, voteId)
	log.Println(err)
	return err == nil
}
