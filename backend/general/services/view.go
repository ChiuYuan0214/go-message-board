package services

import (
	"general/constants"
	"strconv"
)

func RecordView(articleId string) {
	count := cache.HGet(constants.VIEW_COUNT_CACHE_NAME, articleId)
	if count == "" {
		return
	}
	num, err := strconv.ParseInt(count, 10, 64)
	if err != nil {
		return
	}
	cache.HSet(constants.VIEW_COUNT_CACHE_NAME, articleId, strconv.FormatInt(num+1, 10))
}
