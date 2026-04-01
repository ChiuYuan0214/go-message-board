package service

import (
	"general/repo"
	"log"
	"net/http"
)

var _ Vote = (*VoteImpl)(nil)

type VoteImpl struct {
	voteRepo    repo.Vote
	articleRepo repo.Article
	commentRepo repo.Comment
}

func (s *VoteImpl) Run() (err error) {
	return
}

func (s *VoteImpl) Stop() {}

func (s *VoteImpl) Vote(userId, sourceId uint64, score int8, voteType *string) (string, uint64) {
	if *voteType == "article" {
		if _, err := s.articleRepo.GetArticleById(sourceId); err != nil {
			return "source not exist", http.StatusBadRequest
		}
	} else {
		if _, err := s.commentRepo.GetById(sourceId); err != nil {
			return "source not exist", http.StatusBadRequest
		}
	}

	existing, err := s.voteRepo.GetByUserAndSource(userId, sourceId)
	if err == nil {
		if !s.UpdateVote(userId, existing.VoteId, score) {
			return "failed to execute query", http.StatusInternalServerError
		}
		return "", existing.VoteId
	}

	id, err := s.voteRepo.CreateVote(userId, sourceId, score, *voteType)
	if err != nil {
		log.Println(err)
		return "failed to execute query", http.StatusInternalServerError
	}

	return "", id
}

func (s *VoteImpl) UpdateVote(userId, voteId uint64, score int8) bool {
	vote, err := s.voteRepo.GetById(voteId)
	if err != nil {
		log.Println(err)
		return false
	}
	if userId != vote.UserId {
		return false
	}
	if score == vote.Score {
		score = 0
	}
	if err = s.voteRepo.UpdateScore(voteId, score); err != nil {
		log.Println(err)
		return false
	}
	return true
}
