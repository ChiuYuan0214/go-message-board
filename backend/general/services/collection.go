package services

import (
	"general/entities"
	"general/types"
	"log"
)

func GetCollections(userId uint64, page, size int64) []types.CollectionData {
	data := []types.CollectionData{}
	start := (page - 1) * size
	err := db.Raw(`select a.article_id, a.user_id, a.title, a.content, u.username as author, i.file_name as authorImage, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes where source_id = a.article_id and vote_type = 'article' and user_id = ?), 0) as myScore, 
	true as hasCollec, a.publish_time as publishTime 
	from articles a inner join collections collec on collec.article_id = a.article_id and collec.user_id = ? 
	inner join users u on u.user_id = a.user_id 
	left join images i on i.user_id = u.user_id 
	where a.publish_time <= now() order by a.publish_time desc, article_id desc limit ? offset ? `, userId, userId, size, start).Find(&data).Error
	if err != nil {
		log.Println(err)
		return []types.CollectionData{}
	}
	return data
}

func AddCollection(userId, articleId uint64) bool {
	result := db.Create(&entities.Collection{UserId: userId, ArticleId: articleId})
	if result.Error != nil {
		log.Println(result.Error)
	}

	if count := result.RowsAffected; count == 0 {
		return false
	}
	return true
}

func RemoveCollection(userId, articleId uint64) bool {
	result := db.Where("user_id = ? and article_id = ?", userId, articleId).Delete(&entities.Collection{})
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println(result.Error)
		return false
	}
	return true
}
