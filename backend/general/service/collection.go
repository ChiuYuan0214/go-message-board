package service

import (
	"general/repo"
	"general/types"
	"log"
)

var _ Collection = (*CollectionImpl)(nil)

type CollectionImpl struct {
	collectionRepo repo.Collection
}

func (s *CollectionImpl) Run() (err error) {
	return
}

func (s *CollectionImpl) Stop() {}

func (s *CollectionImpl) GetCollections(userId uint64, page, size int64) (data []types.CollectionData) {
	start := (page - 1) * size

	data, err := s.collectionRepo.GetByUserId(userId, start, size)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (s *CollectionImpl) AddCollection(userId, articleId uint64) bool {
	err := s.collectionRepo.Create(userId, articleId)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (s *CollectionImpl) RemoveCollection(userId, articleId uint64) bool {
	err := s.collectionRepo.DeleteByUserIdAndArticleId(userId, articleId)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
