package service

import (
	"general/constants"
	"general/infra"
	"general/repo"
	"general/types"
	"log"
	"strconv"
)

var _ Articles = (*ArticlesImpl)(nil)

type ArticlesImpl struct {
	articleRepo repo.Article
	tagRepo     repo.Tag
	cache       infra.Cache
}

func (s *ArticlesImpl) Run() (err error) {
	return
}

func (s *ArticlesImpl) Stop() {}

// get by newest (default)
func (s *ArticlesImpl) GetNewestList(page, size int64, userId uint64) (list []types.ArticleListData) {
	data := []types.ArticleListData{}
	start := (page - 1) * size
	list, err := s.articleRepo.GetNewsList(userId, start, size)
	if err != nil {
		log.Println(err)
	}

	return data
}

func (s *ArticlesImpl) GetViewList(page, size int64, userId uint64) (data []types.ArticleListData) {
	articleIds := s.cache.LRange(constants.VIEW_LIST_NAME, page, size)
	if len(articleIds) == 0 || page < 1 || size < 1 {
		return data
	}

	data, err := s.articleRepo.GetViewList(userId, articleIds)
	if err != nil {
		log.Println(err)
	}

	return s.sortByOrder(data, articleIds)
}

func (s *ArticlesImpl) GetHotList(page, size int64, userId uint64) (data []types.ArticleListData) {
	articleIds := s.cache.LRange(constants.HOT_LIST_NAMAE, page, size)
	if len(articleIds) == 0 || page < 1 || size < 1 {
		return data
	}

	data, err := s.articleRepo.GetHotList(userId, articleIds)
	if err != nil {
		log.Println(err)
	}

	return s.sortByOrder(data, articleIds)
}

func (s *ArticlesImpl) GetProfileList(page, size int64, userId uint64, selfUserId uint64) (data []types.ArticleListData) {
	start := (page - 1) * size

	data, err := s.articleRepo.GetProfileList(userId, selfUserId, start, size)
	if err != nil {
		log.Println(err)
	}

	return data
}

func (s *ArticlesImpl) GetTagList(page, size int64, tag string) (data []types.ArticleListData) {
	data, err := s.articleRepo.GetTagList(tag)
	if err != nil {
		log.Println(err)
	}

	return data
}

// set tags of each article
func (s *ArticlesImpl) setTags(data []types.ArticleListData, idList []string) {
	if len(idList) == 0 {
		return
	}

	tags, err := s.tagRepo.GetTagsByArticleIds(idList)
	if err != nil {
		log.Println(err)
		return
	}

	articleTagsMap := map[uint64][]string{}
	for _, tag := range tags {
		articleTagsMap[tag.ArticleId] = append(articleTagsMap[tag.ArticleId], tag.Name)
	}

	for _, art := range data {
		tagNames, exist := articleTagsMap[art.ArticleId]
		if exist {
			art.Tags = tagNames
		}
	}
}

func (s *ArticlesImpl) sortByOrder(data []types.ArticleListData, orderList []string) []types.ArticleListData {
	orderMap := map[uint64]int{}

	for i, v := range orderList {
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			continue
		}
		orderMap[id] = i
	}

	sortedData := make([]types.ArticleListData, len(orderList))
	for _, a := range data {
		index := orderMap[a.ArticleId]
		sortedData[index] = a
	}

	compressed := []types.ArticleListData{}
	for _, d := range sortedData {
		if d.ArticleId != 0 {
			compressed = append(compressed, d)
		}
	}

	return compressed
}
