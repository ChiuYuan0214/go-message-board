package services

import (
	"database/sql"
	"fmt"
	"general/constants"
	"general/types"
	"log"
	"strconv"
	"strings"
)

var baseStmt = `
    select article_id, user_id, title, content, top_comment_id, 
	coalesce((select count(vote_id) from votes 
    where source_id = article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, update_time 
	from articles where publish_time <= now() `

// get by newest (default)
func GetNewestList(page, size int64) *[]types.ArticleListData {
	data := &[]types.ArticleListData{}
	start := (page - 1) * size
	stmt := baseStmt + `order by publish_time desc limit ? offset ?`
	rows, err := connPool.Query(stmt, size, start)
	if err != nil {
		log.Println("error when getting newest list:", err)
		return data
	}
	return translate(rows, data)
}

func GetViewList(page, size int64) *[]types.ArticleListData {
	list := cache.LRange(constants.VIEW_LIST_NAME, page, size)
	data := &[]types.ArticleListData{}
	if len(list) == 0 || page < 1 || size < 1 {
		return data
	}

	listStr := strings.Join(list, ", ")
	stmt := fmt.Sprintf(baseStmt+`and article_id in (%s)`, listStr)
	rows, err := connPool.Query(stmt)
	if err != nil {
		log.Println("error when getting view list:", err)
		return data
	}
	data = translate(rows, data)
	return sortByOrder(data, list)
}

func GetHotList(page, size int64) *[]types.ArticleListData {
	list := cache.LRange(constants.HOT_LIST_NAMAE, page, size)
	data := &[]types.ArticleListData{}
	if len(list) == 0 || page < 1 || size < 1 {
		return data
	}

	listStr := strings.Join(list, ", ")
	stmt := fmt.Sprintf(baseStmt+`and article_id in (%s)`, listStr)
	rows, err := connPool.Query(stmt)
	if err != nil {
		log.Println("error when getting hot list:", err)
		return data
	}
	data = translate(rows, data)
	return sortByOrder(data, list)
}

func GetProfileList(page, size, userId int64) *[]types.ArticleListData {
	data := &[]types.ArticleListData{}
	rows, err := connPool.Query(baseStmt+`and user_id = ? order by publish_time desc`, userId)
	if err != nil {
		log.Println("error when getting profile list:", err)
		return data
	}
	return translate(rows, data)
}

func GetTagList(page, size int64, tag string) *[]types.ArticleListData {
	data := &[]types.ArticleListData{}
	rows, err := connPool.Query(`
	select a.article_id, a.user_id, a.title, a.content, a.top_comment_id,
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, update_time 
    from articles a 
    inner join article_tag_maps m on m.article_id = a.article_id inner join tags t on m.tag_id = t.tag_id and t.name = ?
	`, tag)
	if err != nil {
		log.Println("error when getting profile list:", err)
		return data
	}
	return translate(rows, data)
}

func translate(rows *sql.Rows, data *[]types.ArticleListData) *[]types.ArticleListData {
	idList := []string{}
	for rows.Next() {
		var article types.ArticleListData
		var topCommentId sql.NullInt64
		err := rows.Scan(
			&article.ArticleId, &article.UserId, &article.Title,
			&article.Content, &topCommentId,
			&article.VoteUp, &article.VoteDown, &article.UpdateTime)
		if err != nil {
			log.Println(err)
		}
		if topCommentId.Valid {
			val, _ := topCommentId.Value()
			article.TopCommentId = val.(uint64)
		}

		article.Tags = []string{}
		*data = append(*data, article)
		idList = append(idList, strconv.FormatUint(article.ArticleId, 10))
	}

	setTags(data, idList)
	return data
}

// set tags of each article
func setTags(data *[]types.ArticleListData, idList []string) {
	if len(idList) == 0 {
		return
	}
	stmt := fmt.Sprintf(`
	select article_id, name from tags 
	inner join article_tag_maps atm 
	on atm.tag_id = tags.tag_id 
	where atm.article_id in (%s)`, strings.Join(idList, ", "))
	rows, err := connPool.Query(stmt)
	if err != nil {
		log.Println("error when get tags:", err)
	}

	articleTagsMap := map[uint64][]string{}
	for rows.Next() {
		var articleId uint64
		var tagName string
		err := rows.Scan(&articleId, &tagName)
		if err != nil {
			log.Println(err)
		}
		articleTagsMap[articleId] = append(articleTagsMap[articleId], tagName)
	}

	for _, art := range *data {
		tagNames, exist := articleTagsMap[art.ArticleId]
		if exist {
			art.Tags = tagNames
		}
	}
}

func sortByOrder(data *[]types.ArticleListData, orderList []string) *[]types.ArticleListData {
	orderMap := map[uint64]int{}
	for i, v := range orderList {
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			continue
		}
		orderMap[id] = i
	}

	sortedData := make([]types.ArticleListData, len(orderList))
	for _, a := range *data {
		index := orderMap[a.ArticleId]
		sortedData[index] = a
	}

	compressed := &[]types.ArticleListData{}
	for _, d := range sortedData {
		if d.ArticleId != 0 {
			*compressed = append(*compressed, d)
		}
	}

	return compressed
}
