package services

import (
	"general/entities"
	"general/types"
	"log"
	"net/http"
)

func Vote(userId, sourceId uint64, score int8, voteType *string) (string, uint64) {
	var count int64
	var err error
	if *voteType == "article" {
		err = db.Model(&types.Article{}).Where("article_id = ?", sourceId).Count(&count).Error
	} else {
		err = db.Model(&entities.Comment{}).Where("comment_id = ?", sourceId).Count(&count).Error
	}
	if err != nil || count == 0 {
		return "source not exist", http.StatusBadRequest
	}

	var voteId uint64
	err = db.Model(&entities.Vote{}).Select("vote_id").Where("user_id = ? and source_id = ?", userId, sourceId).Error
	if err == nil {
		if UpdateVote(userId, voteId, score) {
			return "", voteId
		} else {
			return "failed to execute query", http.StatusInternalServerError
		}
	}

	newVote := entities.Vote{UserId: userId, SourceId: sourceId, Score: score, VoteType: *voteType}
	result := db.Create(&newVote)
	if result.Error != nil {
		log.Println(result.Error)
		return "failed to execute query", http.StatusInternalServerError
	}

	return "", newVote.VoteId
}

func UpdateVote(userId, voteId uint64, score int8) bool {
	vote := entities.Vote{}
	var curVote struct {
		userId uint64 `gorm:"column:user_id"`
		score  int8
	}
	err := db.Model(&vote).Select("user_id", "score").Where("vote_id = ?", voteId).Scan(&curVote).Error
	if userId != curVote.userId {
		return false
	}
	if score == curVote.score {
		score = 0
	}
	vote.VoteId = voteId
	if err = db.Model(&vote).Update("score", score).Error; err != nil {
		log.Println(err)
	}
	return err == nil
}
