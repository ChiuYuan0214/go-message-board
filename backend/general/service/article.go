package service

import (
	"general/enum"
	"general/repo"
	"general/types"
	"log"
	"net/http"
	"time"
)

var _ Article = (*ArticleImpl)(nil)

type ArticleImpl struct {
	articleRepo    repo.Article
	voteRepo       repo.Vote
	commentRepo    repo.Comment
	tagRepo        repo.Tag
	tagMapRepo     repo.TagMap
	collectionRepo repo.Collection
}

func (s *ArticleImpl) Run() (err error) {
	return
}

func (s *ArticleImpl) Stop() {}

func (s *ArticleImpl) GetArticle(userId uint64, articleId string) (article *types.Article, code int) {
	art, err := s.articleRepo.GetArticleDetail(userId, articleId)
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	return &art, 0
}

func (s *ArticleImpl) InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) uint64 {
	id, err := s.articleRepo.InsertArticle(userId, article, publishTime)
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}

func (s *ArticleImpl) UpdateArticle(userId uint64, articleId uint64, data *types.UpdateArticleData) (string, int) {
	article, err := s.articleRepo.GetArticleById(articleId)
	if err != nil {
		return "article not exist.", http.StatusBadRequest
	}

	if userId != article.UserId {
		return "user incorrect.", http.StatusBadRequest
	}

	updateData := map[string]interface{}{}
	if data.Title != "" {
		updateData["title"] = data.Title
	}
	if data.Content != "" {
		updateData["content"] = data.Content
	}

	if err := s.articleRepo.UpdateArticle(articleId, updateData); err != nil {
		return "failed to update.", http.StatusInternalServerError
	}

	return "", 0
}

func (s *ArticleImpl) DeleteArticle(userId uint64, articleId uint64) (string, int) {
	article, err := s.articleRepo.GetArticleById(articleId)
	if err != nil {
		return "something went wrong", http.StatusInternalServerError
	}

	if userId != article.UserId {
		return "user incorrect", http.StatusBadRequest
	}

	err = s.voteRepo.DeleteVoteBySourceId(enum.VoteSourceArticle, articleId)
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = s.commentRepo.DeleteByArticleId(articleId)
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = s.tagMapRepo.DeleteByArticleId(articleId)
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = s.collectionRepo.DeleteByArticleId(articleId)
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = s.articleRepo.DeleteArticle(articleId)
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	return "", 0
}

func (s *ArticleImpl) GetTagsByArticleId(articleId string) (tags []string) {
	tags, err := s.tagRepo.GetTagsByArticleId(articleId)
	if err != nil {
		log.Println(err)
		return []string{}
	}

	return
}

func (s *ArticleImpl) InsertTags(articleId uint64, tags []string) bool {
	results, err := s.tagRepo.GetByNames(tags)
	if err != nil {
		log.Println(err)
		return false
	}

	tagMap := map[string]bool{}
	for _, e := range results {
		tagMap[e.Name] = true
	}

	for _, name := range tags {
		if !tagMap[name] {
			err := s.tagRepo.Create(name)
			if err != nil {
				return false
			}
		}
	}

	for _, tag := range results {
		err = s.tagMapRepo.CreateIgnoringConflict(articleId, tag.TagId)
		if err != nil {
			log.Println(err)
			return false
		}
	}

	return true
}
