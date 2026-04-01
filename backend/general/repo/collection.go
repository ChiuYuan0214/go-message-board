package repo

import (
	"general/entities"
	"general/infra"
	"general/types"
)

var _ Collection = (*CollectionImpl)(nil)

type CollectionImpl struct {
	db infra.RDB
}

func (r *CollectionImpl) Run() (err error) {
	return
}

func (r *CollectionImpl) Stop() {}

func (r *CollectionImpl) Create(userId, articleId uint64) (err error) {
	err = r.db.Orm().Create(
		&entities.Collection{
			UserId:    userId,
			ArticleId: articleId,
		},
	).Error
	return
}

func (r *CollectionImpl) DeleteByUserIdAndArticleId(userId, articleId uint64) (err error) {
	err = r.db.Orm().
		Where("user_id = ? and article_id = ?", userId, articleId).
		Delete(new(entities.Collection)).Error
	return
}

func (r *CollectionImpl) DeleteByArticleId(articleId uint64) (err error) {
	err = r.db.Orm().
		Where("article_id = ?", articleId).
		Delete(new(entities.Collection)).Error
	return
}

func (r *CollectionImpl) GetByUserId(userId uint64, start, size int64) (data []types.CollectionData, err error) {
	err = r.db.Orm().Raw(
		`select 
	a.article_id, 
	a.user_id, 
	a.title, 
	a.content, 
	u.username as author, 
	i.file_name as authorImage, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id 
	and vote_type = 'article' 
	and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id 
	and vote_type = 'article' 
	and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes 
	where source_id = a.article_id 
	and vote_type = 'article' 
	and user_id = ?), 0) as myScore, 
	true as hasCollec, a.publish_time as publishTime 
	from articles a 
	inner join collections collec on collec.article_id = a.article_id and collec.user_id = ? 
	inner join users u on u.user_id = a.user_id 
	left join images i on i.user_id = u.user_id 
	where a.publish_time <= now() 
	order by a.publish_time desc, article_id desc limit ? offset ? `, userId, userId, size, start).
		Find(&data).Error

	return
}
