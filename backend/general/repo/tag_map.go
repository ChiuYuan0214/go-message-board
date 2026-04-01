package repo

import (
	"general/entities"
	"general/infra"

	"gorm.io/gorm/clause"
)

var _ TagMap = (*TagMapImpl)(nil)

type TagMapImpl struct {
	db infra.RDB
}

func (r *TagMapImpl) Run() (err error) {
	return
}

func (r *TagMapImpl) Stop() {}

func (r *TagMapImpl) DeleteByArticleId(articleId uint64) (err error) {
	err = r.db.Orm().
		Where("article_id = ?", articleId).
		Delete(new(entities.ArticleTagMap)).Error
	return
}

func (r *TagMapImpl) CreateIgnoringConflict(articleId uint64, tagId uint64) (err error) {
	err = r.db.Orm().
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(
			&entities.ArticleTagMap{
				ArticleId: articleId,
				TagId:     tagId,
			},
		).Error
	return
}
