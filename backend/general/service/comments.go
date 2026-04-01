package service

import (
	"general/repo"
	"general/types"
	"log"
)

var _ Comments = (*CommentsImpl)(nil)

type CommentsImpl struct {
	commentRepo repo.Comment
}

func (s *CommentsImpl) Run() (err error) {
	return
}

func (s *CommentsImpl) Stop() {}

func (s *CommentsImpl) GetComments(articleId uint64) (data []types.CommentListData) {
	data, err := s.commentRepo.GetByArticleId(articleId)
	if err != nil {
		log.Println("failed to get comment list:", err)
		return
	}

	return
}
