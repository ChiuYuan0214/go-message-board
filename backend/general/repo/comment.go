package repo

import (
	"general/entities"
	"general/infra"
	"general/types"
)

var _ Comment = (*CommentImpl)(nil)

type CommentImpl struct {
	db infra.RDB
}

func (r *CommentImpl) Run() (err error) {
	return
}

func (r *CommentImpl) Stop() {}

func (r *CommentImpl) Create(userId, articleId uint64, title, content string) (newComment *entities.Comment, err error) {
	newComment = &entities.Comment{
		UserId:    userId,
		ArticleId: articleId,
		Title:     title,
		Content:   content,
	}
	err = r.db.Orm().Create(newComment).Error
	return
}

func (r *CommentImpl) Updates(commentId uint64, fields map[string]any) (err error) {
	err = r.db.Orm().
		Model(new(entities.Comment)).
		Where("comment_id = ?", commentId).
		Updates(fields).Error
	return
}

func (r *CommentImpl) DeleteById(commentId uint64) (err error) {
	err = r.db.Orm().
		Where("comment_id = ?", commentId).
		Delete(new(entities.Comment)).Error
	return
}

func (r *CommentImpl) DeleteByArticleId(articleId uint64) (err error) {
	err = r.db.Orm().
		Where("article_id = ?", articleId).
		Delete(new(entities.Comment)).Error
	return
}

func (r *CommentImpl) GetById(commentId uint64) (comment entities.Comment, err error) {
	err = r.db.Orm().
		Where("comment_id = ?", commentId).
		Take(&comment).Error
	return
}

func (r *CommentImpl) GetByArticleId(articleId uint64) (data []types.CommentListData, err error) {
	err = r.db.Orm().Raw(`select c.comment_id, c.user_id, u.username as commenter, i.file_name as commenterImage, 
	c.title, c.content, c.creation_time as creationTime, 
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = 1), 0) as voteUp,
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = -1), 0) as voteDown 
    from comments c inner join users u on u.user_id = c.user_id left join images i on i.user_id = u.user_id 
	where c.article_id = ? order by c.creation_time desc`, articleId).Find(&data).Error
	return
}
