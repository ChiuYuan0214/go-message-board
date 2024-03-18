package services

import (
	"general/entities"
	"general/types"
	"log"
	"net/http"
	"time"
)

func GetArticle(userId uint64, articleId string) (*types.Article, int) {
	var article types.Article
	err := db.Raw(`
	select a.article_id, a.user_id, u.username, i.file_name, a.title, a.content, a.top_comment_id, a.edited, a.view_count, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes where source_id = a.article_id and vote_type = 'article' and user_id = ?), 0) as myScore, 
	((select collec.article_id from collections collec where collec.article_id = a.article_id and collec.user_id = ?) is not null) as hasCollec, 
	a.publish_time, a.creation_time, a.update_time 
	from articles a inner join users u on u.user_id = a.user_id left join images i on i.user_id = u.user_id 
	where a.article_id = ?`, userId, userId, articleId).Scan(&article).Error
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	return &article, 0
}

func InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) uint64 {
	log.Println("userId in InsertArticle:", userId)
	newArticle := types.Article{
		UserId:      userId,
		Title:       article.Title,
		Content:     article.Content,
		PublishTime: *publishTime,
	}
	err := db.Create(&newArticle).Error
	if err != nil {
		log.Println(err)
		return 0
	}
	return newArticle.ArticleId
}

func UpdateArticle(userId uint64, articleId uint64, data *types.UpdateArticleData) (string, int) {
	var article types.Article
	if err := db.Where("article_id = ?", articleId).First(&article).Error; err != nil {
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

	if err := db.Model(&article).Updates(updateData).Error; err != nil {
		return "failed to update.", http.StatusInternalServerError
	}

	return "", 0
}

func DeleteArticle(userId uint64, articleId uint64) (string, int) {
	var actualUserId uint64
	err := db.Table("articles").Select("user_id").Where("article_id = ?", articleId).Scan(&actualUserId).Error
	if err != nil {
		return "something went wrong", http.StatusInternalServerError
	}
	if userId != actualUserId {
		return "user incorrect", http.StatusBadRequest
	}

	tags := []string{}
	DeleteRemovedTags(articleId, tags)

	err = db.Where("vote_type = 'article' and source_id = ?", articleId).Delete(&entities.Vote{}).Error
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = db.Where("article_id = ?", articleId).Delete(&entities.Comment{}).Error
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = db.Where("article_id = ?", articleId).Delete(&entities.ArticleTagMap{}).Error
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = db.Where("article_id = ?", articleId).Delete(&entities.Collection{}).Error
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	err = db.Delete(&types.Article{ArticleId: articleId}).Error
	if err != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	return "", 0
}

func GetTagsByArticleId(articleId string) []string {
	tagList := []string{}
	err := db.Model(&entities.Tag{}).
		Joins("inner join article_tag_maps atm on atm.tag_id = tags.tag_id where atm.article_id = ?", articleId).
		Pluck("name", &tagList).Error
	if err != nil {
		return []string{}
	}
	return tagList
}

func InsertTags(articleId uint64, tags []string) bool {
	var results []struct {
		id   int64
		name string
	}
	err := db.Model(&entities.Tag{}).Select("tag_id", "name").Where("name in (?)", tags).Find(&results).Error
	if err != nil {
		log.Println(err)
		return false
	}

	tagMap := map[string]bool{}
	for _, e := range results {
		tagMap[e.name] = true
	}

	for _, name := range tags {
		if !tagMap[name] {
			err := db.Create(&entities.Tag{Name: name}).Error
			if err != nil {
				return false
			}
		}
	}

	tagIdList := []uint64{}
	err = db.Model(&entities.Tag{}).Select("tag_id").Where("name in (?)", tags).
		Where("tag_id not in (?)", db.Table("article_tag_maps").
			Select("tag_id").Where("article_id = ?", articleId)).Find(&tagIdList).Error
	if err != nil {
		return false
	}

	for _, tagId := range tagIdList {
		err = db.Create(&entities.ArticleTagMap{ArticleId: articleId, TagId: tagId}).Error
		if err != nil {
			log.Println(err)
			return false
		}
	}

	return true
}

func DeleteRemovedTags(articleId uint64, tags []string) bool {
	tagMap := map[string]bool{}
	for _, name := range tags {
		tagMap[name] = true
	}

	var deleteList []string
	err := db.Model(&entities.Tag{}).Where("tag_id in (?)",
		db.Table("article_tag_maps").Select("tag_id").Where("article_id = ?", articleId)).
		Scan(&deleteList).Error
	if err != nil {
		return false
	}
	if len(deleteList) == 0 {
		return true
	}

	err = db.Where("tag_id in (?)", db.Table("tags").
		Where("name in (?)", deleteList)).Delete(&entities.ArticleTagMap{}).Error
	return err == nil
}
