package services

import (
	"database/sql"
	"general/types"
	"log"
)

func GetCollections(userId int64, page, size int64) *[]types.CollectionData {
	data := &[]types.CollectionData{}
	start := (page - 1) * size
	rows, err := connPool.Query(`select a.article_id, a.user_id, a.title, a.content, u.username, i.file_name, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes where source_id = a.article_id and vote_type = 'article' and user_id = ?), 0) as myScore, 
	true as hasCollec, a.publish_time 
	from articles a inner join collections collec on collec.article_id = a.article_id and collec.user_id = ? 
	inner join users u on u.user_id = a.user_id 
	left join images i on i.user_id = u.user_id 
	where a.publish_time <= now() order by publish_time desc, article_id desc limit ? offset ? `, userId, userId, size, start)
	if err != nil {
		log.Println(err)
		return data
	}
	for rows.Next() {
		var collec types.CollectionData
		var authorImage sql.NullString
		err := rows.Scan(&collec.ArticleId, &collec.UserId, &collec.Title, &collec.Content, &collec.Author,
			&authorImage, &collec.VoteUp, &collec.VoteDown, &collec.MyScore, &collec.HasCollec, &collec.PublishTime)
		if err != nil {
			log.Println(err)
			continue
		}
		if authorImage.Valid {
			val, _ := authorImage.Value()
			collec.AuthorImage = val.(string)
		}
		*data = append(*data, collec)
	}
	return data
}

func AddCollection(userId int64, articleId uint64) bool {
	sqlRes, err := connPool.Exec(`insert into collections (user_id, article_id) values (?, ?)`, userId, articleId)
	if err != nil {
		log.Println(err)
	}
	if count, err2 := sqlRes.RowsAffected(); err != nil || count == 0 || err2 != nil {
		log.Println(err)
		return false
	}
	return true
}

func RemoveCollection(userId int64, articleId uint64) bool {
	sqlRes, err := connPool.Exec(`delete from collections where user_id = ? and article_id = ?`, userId, articleId)
	if count, err2 := sqlRes.RowsAffected(); err != nil || count == 0 || err2 != nil {
		log.Println(err)
		return false
	}
	return true
}
