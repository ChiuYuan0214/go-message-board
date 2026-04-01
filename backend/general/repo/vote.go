package repo

import (
	"general/entities"
	"general/enum"
	"general/infra"
)

var _ Vote = (*VoteImpl)(nil)

type VoteImpl struct {
	db infra.RDB
}

func (r *VoteImpl) Run() (err error) {
	return
}

func (r *VoteImpl) Stop() {}

func (r *VoteImpl) GetByUserAndSource(userId, sourceId uint64) (vote entities.Vote, err error) {
	err = r.db.Orm().
		Where("user_id = ? AND source_id = ?", userId, sourceId).
		First(&vote).Error
	return
}

func (r *VoteImpl) CreateVote(userId, sourceId uint64, score int8, voteType string) (id uint64, err error) {
	vote := entities.Vote{UserId: userId, SourceId: sourceId, Score: score, VoteType: voteType}
	err = r.db.Orm().Create(&vote).Error
	id = vote.VoteId
	return
}

func (r *VoteImpl) GetById(voteId uint64) (vote entities.Vote, err error) {
	err = r.db.Orm().
		Where("vote_id = ?", voteId).
		First(&vote).Error
	return
}

func (r *VoteImpl) UpdateScore(voteId uint64, score int8) (err error) {
	err = r.db.Orm().
		Model(&entities.Vote{VoteId: voteId}).
		Update("score", score).Error
	return
}

func (r *VoteImpl) DeleteVoteBySourceId(source enum.VoteSource, sourceId uint64) (err error) {
	err = r.db.Orm().
		Where("vote_type = ?", source).
		Where("source_id = ?", sourceId).
		Delete(new(entities.Vote)).Error
	return
}
