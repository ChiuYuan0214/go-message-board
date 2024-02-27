package services

import (
	"database/sql"
	"fmt"
	"general/types"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetArticle(userId int64, articleId string) (*types.Article, int) {
	var article types.Article
	rows, err := connPool.Query(`
	select a.article_id, a.user_id, u.username, i.file_name, a.title, a.content, a.top_comment_id, a.edited, a.view_count, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes where source_id = a.article_id and vote_type = 'article' and user_id = ?), 0) as myScore, 
	((select collec.article_id from collections collec where collec.article_id = a.article_id and collec.user_id = ?) is not null) as hasCollec, 
	a.publish_time, a.creation_time, a.update_time 
	from articles a inner join users u on u.user_id = a.user_id left join images i on i.user_id = u.user_id 
	where a.article_id = ?`, userId, userId, articleId)
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError
	}
	defer rows.Close()

	var topCommentId sql.NullInt64
	var viewCnt sql.NullInt64
	var publishTime sql.NullTime
	var authorImage sql.NullString

	exist := rows.Next()
	if !exist {
		return nil, http.StatusBadRequest
	}

	err = rows.Scan(&article.ArticleId, &article.UserId, &article.Author, &authorImage, &article.Title, &article.Content,
		&topCommentId, &article.Edited, &viewCnt, &article.VoteUp, &article.VoteDown, &article.MyScore, &article.HasCollec,
		&publishTime, &article.CreationTime, &article.UpdateTime)
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	if topCommentId.Valid {
		article.TopCommentId = uint64(topCommentId.Int64)
	}
	if viewCnt.Valid {
		article.ViewCount = uint32(viewCnt.Int64)
	}
	if publishTime.Valid {
		article.PublishTime = publishTime.Time.Local()
	}
	if authorImage.Valid {
		article.AuthorImage = authorImage.String
	}

	return &article, 0
}

func InsertArticle(userId int64, article *types.AddArticleData, publishTime *time.Time) int64 {
	// parsedTime, err := time.Parse("2006-01-02 15:04:05", article.PublishTime)
	sqlRes, err := connPool.Exec("insert into articles (user_id, title, content, publish_time) values (?, ?, ?, ?)",
		userId, article.Title, article.Content, publishTime)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := sqlRes.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}

	return id
}

func UpdateArticle(userId int64, articleId int64, data *types.UpdateArticleData) (string, int) {
	row := connPool.QueryRow("select user_id from articles where article_id = ?", articleId)
	var actualUserId int64
	err := row.Scan(&actualUserId)
	if err != nil {
		return "article not exist.", http.StatusBadRequest
	}

	if userId != actualUserId {
		return "user incorrect.", http.StatusBadRequest
	}

	updateList := []string{}
	if data.Title != "" {
		updateList = append(updateList, fmt.Sprintf(" title = '%s'", data.Title))
	}
	if data.Content != "" {
		updateList = append(updateList, fmt.Sprintf(" content = '%s'", data.Content))
	}

	stmt := "update articles set edited = true, " + strings.Join(updateList, ", ") + " where article_id = ?"
	_, err = connPool.Exec(stmt, articleId)

	if err != nil {
		return "failed to update.", http.StatusInternalServerError
	}

	return "", 0
}

func DeleteArticle(userId int64, articleId int64) (string, int) {
	var actualUserId int64
	row := connPool.QueryRow("select user_id from articles where article_id = ?", articleId)
	err := row.Scan(&actualUserId)
	if err != nil {
		return "something went wrong", http.StatusInternalServerError
	}
	if userId != actualUserId {
		return "user incorrect", http.StatusBadRequest
	}

	tags := []string{}
	DeleteRemovedTags(articleId, &tags)

	sqlRes, err := connPool.Exec("delete from votes where vote_type = 'article' and source_id = ?", articleId)
	_, err2 := sqlRes.RowsAffected()
	if err != nil || err2 != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	sqlRes, err = connPool.Exec("delete from comments where article_id = ?", articleId)
	_, err2 = sqlRes.RowsAffected()
	if err != nil || err2 != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	sqlRes, err = connPool.Exec("delete from article_tag_maps where article_id = ?", articleId)
	_, err2 = sqlRes.RowsAffected()
	if err != nil || err2 != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	sqlRes, err = connPool.Exec("delete from collections where article_id = ?", articleId)
	_, err2 = sqlRes.RowsAffected()
	if err != nil || err2 != nil {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	sqlRes, err = connPool.Exec("delete from articles where article_id = ?", articleId)
	count, _ := sqlRes.RowsAffected()
	if err != nil || count < 1 {
		log.Println(err)
		return "failed to delete article.", http.StatusInternalServerError
	}

	return "", 0
}

func GetTagsByArticleId(articleId string) []string {
	tagList := []string{}
	rows, err := connPool.Query("select name from tags inner join article_tag_maps atm on atm.tag_id = tags.tag_id where atm.article_id = ?", articleId)
	if err != nil {
		return tagList
	}
	defer rows.Close()

	for rows.Next() {
		var tagName string
		err := rows.Scan(&tagName)
		if err != nil {
			continue
		}
		tagList = append(tagList, tagName)
	}
	return tagList
}

func InsertTags(articleId int64, tags *([]string)) bool {
	stmt := fmt.Sprintf("select tag_id, name from tags where name in (%s)", "'"+strings.Join(*tags, "', '")+"'")
	rows, err := connPool.Query(stmt)
	if err != nil {
		return false
	}
	defer rows.Close()

	tagMap := map[string]bool{}
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return false
		}
		tagMap[name] = true
	}

	for _, name := range *tags {
		if !tagMap[name] {
			_, err := connPool.Exec("insert into tags (name) values (?)", name)
			if err != nil {
				return false
			}
		}
	}

	stmt = fmt.Sprintf(`select tag_id from tags where name in (%s) 
	and tag_id not in (select tag_id from article_tag_maps where article_id = %s)`,
		"'"+strings.Join(*tags, "', '")+"'", strconv.FormatInt(articleId, 10))
	rows, err = connPool.Query(stmt)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var tagId int64
		err := rows.Scan(&tagId)
		if err != nil {
			return false
		}
		_, err = connPool.Exec("insert into article_tag_maps (article_id, tag_id) values (?, ?)", articleId, tagId)
		if err != nil {
			return false
		}
	}

	return true
}

func DeleteRemovedTags(articleId int64, tags *([]string)) bool {
	tagMap := map[string]bool{}
	for _, name := range *tags {
		tagMap[name] = true
	}

	var deleteList []string
	rows, err := connPool.Query("select name from tags where tag_id in (select tag_id from article_tag_maps where article_id = ?)", articleId)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return false
		}
		if !tagMap[name] {
			deleteList = append(deleteList, name)
		}
	}

	stmt := fmt.Sprintf("delete from article_tag_maps where tag_id in (select tag_id from tags where name in (%s))", "'"+strings.Join(deleteList, "', '")+"'")
	_, err = connPool.Exec(stmt)
	return err == nil
}
