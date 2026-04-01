package service

import (
	"general/constants"
	"general/enum"
	"general/infra"
	"general/repo"
	"general/types"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var _ Comment = (*CommentImpl)(nil)

type CommentImpl struct {
	articleRepo repo.Article
	commentRepo repo.Comment
	voteRepo    repo.Vote
	cache       infra.Cache
}

func (s *CommentImpl) Run() (err error) {
	return
}

func (s *CommentImpl) Stop() {}

func (s *CommentImpl) AddComment(userId uint64, data *types.AddCommentData) uint64 {
	_, err := s.articleRepo.GetArticleById(data.ArticleId)
	if err != nil {
		log.Println(err)
		return 0
	}

	newComment, err := s.commentRepo.Create(userId, data.ArticleId, data.Title, data.Content)
	if err != nil {
		log.Println(err)
		return 0
	}

	s.recordArticle(data.ArticleId)

	return newComment.CommentId
}

func (s *CommentImpl) UpdateComment(userId uint64, data *types.UpdateCommentData) (string, int) {
	comment, err := s.commentRepo.GetById(data.CommentId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	if userId != comment.UserId {
		return "user incorrect.", http.StatusBadRequest
	}

	fields := map[string]any{}
	if strings.Trim(data.Title, " ") != "" {
		fields["title"] = data.Title
	}
	if strings.Trim(data.Content, " ") != "" {
		fields["content"] = data.Content
	}
	if len(fields) == 0 {
		return "nothing to update.", http.StatusBadRequest
	}

	s.commentRepo.Updates(data.CommentId, fields)
	if err != nil {
		return "something went wrong", http.StatusInternalServerError
	}

	return "", 0
}

func (s *CommentImpl) DeleteComment(userId uint64, commentIdStr string) (string, int) {
	id, _ := strconv.Atoi(commentIdStr)
	commentId := uint64(id)
	comment, err := s.commentRepo.GetById(commentId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}
	if userId != comment.UserId {
		return "user incorrect.", http.StatusBadRequest
	}

	// delete related votes
	err = s.voteRepo.DeleteVoteBySourceId(enum.VoteSourcceComment, commentId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	// delete comment
	s.commentRepo.DeleteById(commentId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	return "", 0
}

func (s *CommentImpl) recordArticle(articleId uint64) error {
	id := strconv.FormatUint(articleId, 10)
	return s.cache.SAdd(constants.COMMENTED_ARTICLE_SET, id)
}
