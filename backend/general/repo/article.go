package repo

import (
	"fmt"
	"general/infra"
	"general/static"
	"general/types"
	"strings"
	"time"
)

var _ Article = (*ArticleImpl)(nil)

type ArticleImpl struct {
	db infra.RDB
}

func (r *ArticleImpl) Run() (err error) {
	return
}

func (r *ArticleImpl) Stop() {}

func (r *ArticleImpl) InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) (id uint64, err error) {
	newArticle := types.Article{
		UserId:      userId,
		Title:       article.Title,
		Content:     article.Content,
		PublishTime: *publishTime,
	}
	err = r.db.Orm().Create(&newArticle).Error
	id = newArticle.ArticleId
	return
}

func (r *ArticleImpl) UpdateArticle(articleId uint64, data map[string]interface{}) (err error) {
	err = r.db.Orm().Where("article_id = ?", articleId).Updates(data).Error
	return
}

func (r *ArticleImpl) DeleteArticle(articleId uint64) (err error) {
	err = r.db.Orm().Delete(&types.Article{ArticleId: articleId}).Error
	return
}

func (r *ArticleImpl) GetArticleById(articleId uint64) (article types.Article, err error) {
	err = r.db.Orm().Where("article_id = ?", articleId).First(&article).Error

	return
}

func (r *ArticleImpl) GetArticleDetail(userId uint64, articleId string) (article types.Article, err error) {
	err = r.db.Orm().Raw(`
	select 
	a.article_id, 
	a.user_id, u.username, 
	i.file_name, 
	a.title, 
	a.content, 
	a.top_comment_id,
	a.edited, 
	a.view_count, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id 
	and vote_type = 'article' 
	and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id 
	and vote_type = 'article' 
	and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes where source_id = a.article_id and vote_type = 'article' and user_id = ?), 0) as myScore, 
	((select collec.article_id from collections collec where collec.article_id = a.article_id and collec.user_id = ?) is not null) as hasCollec, 
	a.publish_time, a.creation_time, a.update_time 
	from articles a 
	inner join users u on u.user_id = a.user_id 
	left join images i on i.user_id = u.user_id 
	where a.article_id = ?`, userId, userId, articleId).Scan(&article).Error

	return
}

func (r *ArticleImpl) GetNewsList(userId uint64, start, size int64) (list []types.ArticleListData, err error) {
	stmt := static.BaseSearchArticlesStmt + `order by a.publish_time desc, a.article_id desc limit ? offset ?`
	err = r.db.Orm().Raw(stmt, userId, userId, size, start).Find(&list).Error
	return
}

func (r *ArticleImpl) GetViewList(userId uint64, articleIds []string) (list []types.ArticleListData, err error) {
	articleIdsStr := strings.Join(articleIds, ", ")
	stmt := fmt.Sprintf(static.BaseSearchArticlesStmt+`and a.article_id in (%s)`, articleIdsStr)
	err = r.db.Orm().Raw(stmt, userId, userId).Find(&list).Error
	return
}

func (r *ArticleImpl) GetHotList(userId uint64, articleIds []string) (list []types.ArticleListData, err error) {
	articleIdsStr := strings.Join(articleIds, ", ")
	stmt := fmt.Sprintf(static.BaseSearchArticlesStmt+`and a.article_id in (%s)`, articleIdsStr)
	err = r.db.Orm().Raw(stmt, userId, userId).Find(&list).Error
	return
}

func (r *ArticleImpl) GetProfileList(userId, selfUserId uint64, start, size int64) (list []types.ArticleListData, err error) {
	err = r.db.Orm().Raw(static.BaseSearchArticlesStmt+
		`and a.user_id = ? order by a.publish_time desc, 
		a.article_id desc limit ? offset ?`, selfUserId, selfUserId, userId, size, start).
		Scan(&list).Error

	return
}

func (r *ArticleImpl) GetTagList(tag string) (list []types.ArticleListData, err error) {
	err = r.db.Orm().Raw(`
	select a.article_id, a.user_id, a.title, a.content, a.top_comment_id,
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, update_time 
    from articles a 
    inner join article_tag_maps m on m.article_id = a.article_id inner join tags t on m.tag_id = t.tag_id and t.name = ?
	`, tag).Find(&list).Error
	return
}
