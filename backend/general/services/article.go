package services

import (
	"database/sql"
	"fmt"
	"general/types"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetArticle(articleId string) (*types.Article, int) {
	var article types.Article
	rows, err := connPool.Query(`
	select a.article_id, 
    (select sum(score) from votes v where v.source_id = a.article_id and v.vote_type = 'article') 
    as score, count(distinct c.comment_id) as commentCount from articles a 
    left join comments c on c.article_id = a.article_id 
    where a.publish_time <= now() 
    group by a.article_id order by a.publish_time desc;`, articleId)
	defer rows.Close()

	if err != nil {
		return nil, http.StatusInternalServerError
	}

	var topCommentId sql.NullInt64
	var viewCnt sql.NullInt64
	var publishTime sql.NullTime

	exist := rows.Next()
	if !exist {
		return nil, http.StatusBadRequest
	}

	err = rows.Scan(&article.ArticleId, &article.UserId, &article.Title, &article.Content,
		&topCommentId, &article.Edited, &viewCnt, &article.VoteUp, &article.VoteDown, &publishTime,
		&article.CreationTime, &article.UpdateTime)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	if topCommentId.Valid {
		article.TopCommentId = uint64(topCommentId.Int64)
	} else {
		article.TopCommentId = 0
	}
	if viewCnt.Valid {
		article.ViewCount = uint32(viewCnt.Int64)
	}
	if publishTime.Valid {
		article.PublishTime = publishTime.Time.Local()
	} else {
		article.PublishTime = time.Time{}
	}

	return &article, 0
}

func InsertArticle(userId int64, article *types.AddArticleData) int64 {
	// parsedTime, err := time.Parse("2006-01-02 15:04:05", article.PublishTime)
	sqlRes, err := connPool.Exec("insert into articles (user_id, title, content, publish_time) values (?, ?, ?, ?)", userId, article.Title, article.Content, article.PublishTime)
	if err != nil {
		return 0
	}
	id, err := sqlRes.LastInsertId()
	if err != nil {
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

	sqlRes, err := connPool.Exec("delete from articles where article_id = ?", articleId)
	count, err := sqlRes.RowsAffected()
	if err != nil || count < 1 {
		return "failed to delete article.", http.StatusInternalServerError
	}

	return "", 0
}

func GetTagsByArticleId(articleId string) []string {
	tagList := []string{}
	rows, err := connPool.Query("select name from tags inner join article_tag_maps atm on atm.tag_id = tags.tag_id where atm.article_id = ?", articleId)
	defer rows.Close()
	if err != nil {
		return tagList
	}

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
	defer rows.Close()
	if err != nil {
		return false
	}

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
	defer rows.Close()
	if err != nil {
		return false
	}

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
	defer rows.Close()
	if err != nil {
		return false
	}

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
	if err != nil {
		return false
	}

	return true
}
