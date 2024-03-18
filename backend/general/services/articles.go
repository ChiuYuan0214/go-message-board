package services

import (
	"fmt"
	"general/constants"
	"general/types"
	"log"
	"strconv"
	"strings"
)

var baseStmt = `
    select a.article_id, a.user_id, a.title, a.content, u.username, i.file_name, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes where source_id = a.article_id and vote_type = 'article' and user_id = ?), 0) as myScore, 
	((select a.article_id from collections collec where collec.article_id = a.article_id and collec.user_id = ?) is not null) as hasCollec, 
	c.title, c.content, cu.username, ci.file_name, a.publish_time 
	from articles a 
	inner join users u on u.user_id = a.user_id 
	left join comments c on top_comment_id = comment_id 
	left join users cu on cu.user_id = c.user_id 
	left join images i on i.user_id = u.user_id 
	left join images ci on ci.user_id = c.user_id 
	where a.publish_time <= now() `

// get by newest (default)
func GetNewestList(page, size int64, userId uint64) []types.ArticleListData {
	data := []types.ArticleListData{}
	start := (page - 1) * size
	stmt := baseStmt + `order by a.publish_time desc, a.article_id desc limit ? offset ?`
	err := db.Raw(stmt, userId, userId, size, start).Find(&data).Error
	if err != nil {
		log.Println(err)
	}

	return data
}

func GetViewList(page, size int64, userId uint64) []types.ArticleListData {
	list := cache.LRange(constants.VIEW_LIST_NAME, page, size)
	data := []types.ArticleListData{}
	if len(list) == 0 || page < 1 || size < 1 {
		return data
	}

	listStr := strings.Join(list, ", ")
	stmt := fmt.Sprintf(baseStmt+`and a.article_id in (%s)`, listStr)
	err := db.Raw(stmt, userId, userId).Find(&data).Error
	if err != nil {
		log.Println(err)
	}
	return sortByOrder(data, list)
}

func GetHotList(page, size int64, userId uint64) []types.ArticleListData {
	list := cache.LRange(constants.HOT_LIST_NAMAE, page, size)
	data := []types.ArticleListData{}
	if len(list) == 0 || page < 1 || size < 1 {
		return data
	}
	listStr := strings.Join(list, ", ")
	stmt := fmt.Sprintf(baseStmt+`and a.article_id in (%s)`, listStr)
	err := db.Raw(stmt, userId, userId).Find(&data).Error
	if err != nil {
		log.Println(err)
	}

	return sortByOrder(data, list)
}

func GetProfileList(page, size int64, userId uint64, selfUserId uint64) []types.ArticleListData {
	data := []types.ArticleListData{}
	start := (page - 1) * size
	err := db.Raw(baseStmt+
		`and a.user_id = ? order by a.publish_time desc, 
		a.article_id desc limit ? offset ?`, selfUserId, selfUserId, userId, size, start).
		Scan(&data).Error
	if err != nil {
		log.Println(err)
	}
	return data
}

func GetTagList(page, size int64, tag string) []types.ArticleListData {
	data := []types.ArticleListData{}
	err := db.Raw(`
	select a.article_id, a.user_id, a.title, a.content, a.top_comment_id,
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, update_time 
    from articles a 
    inner join article_tag_maps m on m.article_id = a.article_id inner join tags t on m.tag_id = t.tag_id and t.name = ?
	`, tag).Find(&data).Error
	if err != nil {
		log.Println(err)
	}
	return data
}

// set tags of each article
func setTags(data []types.ArticleListData, idList []string) {
	if len(idList) == 0 {
		return
	}
	var tags []struct {
		ArticleId uint64
		Name      string
	}

	err := db.Table("tags").
		Select("article_id, name").
		Joins("INNER JOIN article_tag_maps ON article_tag_maps.tag_id = tags.tag_id").
		Where("article_tag_maps.article_id IN (?)", idList).
		Scan(&tags).Error
	if err != nil {
		log.Println(err)
		return
	}

	articleTagsMap := map[uint64][]string{}
	for _, tag := range tags {
		articleTagsMap[tag.ArticleId] = append(articleTagsMap[tag.ArticleId], tag.Name)
	}

	for _, art := range data {
		tagNames, exist := articleTagsMap[art.ArticleId]
		if exist {
			art.Tags = tagNames
		}
	}
}

func sortByOrder(data []types.ArticleListData, orderList []string) []types.ArticleListData {
	orderMap := map[uint64]int{}
	for i, v := range orderList {
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			continue
		}
		orderMap[id] = i
	}

	sortedData := make([]types.ArticleListData, len(orderList))
	for _, a := range data {
		index := orderMap[a.ArticleId]
		sortedData[index] = a
	}

	compressed := []types.ArticleListData{}
	for _, d := range sortedData {
		if d.ArticleId != 0 {
			compressed = append(compressed, d)
		}
	}

	return compressed
}
