package service

import (
	"general/constants"
	"general/infra"
	"strconv"
)

var _ View = (*ViewImpl)(nil)

type ViewImpl struct {
	cache infra.Cache
}

func (s *ViewImpl) Run() (err error) {
	return
}

func (s *ViewImpl) Stop() {}

func (s *ViewImpl) RecordView(articleId string) {
	count := s.cache.HGet(constants.VIEW_COUNT_CACHE_NAME, articleId)
	if count == "" {
		return
	}

	num, err := strconv.ParseInt(count, 10, 64)
	if err != nil {
		return
	}

	s.cache.HSet(constants.VIEW_COUNT_CACHE_NAME, articleId, strconv.FormatInt(num+1, 10))
}
