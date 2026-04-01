package repo

import (
	"general/dto"
	"general/entities"
	"general/infra"
)

var _ Tag = (*TagImpl)(nil)

type TagImpl struct {
	db infra.RDB
}

func (r *TagImpl) Run() (err error) {
	return
}

func (r *TagImpl) Stop() {}

func (r *TagImpl) Create(name string) (err error) {
	err = r.db.Orm().Create(&entities.Tag{Name: name}).Error
	return
}

func (r *TagImpl) GetTagsByArticleId(articleId string) (tags []string, err error) {
	err = r.db.Orm().
		Model(new(entities.Tag)).
		Joins("inner join article_tag_maps atm on atm.tag_id = tags.tag_id where atm.article_id = ?", articleId).
		Pluck("name", &tags).Error
	return
}

func (r *TagImpl) GetByNames(names []string) (tags []entities.Tag, err error) {
	err = r.db.Orm().Where("name in (?)", tags).Find(&tags).Error
	return
}

func (r *TagImpl) GetTagsByArticleIds(articleIds []string) (data []dto.ArticleTag, err error) {
	err = r.db.Orm().Table("tags").
		Select("article_id, name").
		Joins("INNER JOIN article_tag_maps ON article_tag_maps.tag_id = tags.tag_id").
		Where("article_tag_maps.article_id IN (?)", articleIds).
		Scan(&data).Error

	return
}
