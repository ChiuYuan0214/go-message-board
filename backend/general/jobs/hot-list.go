package jobs

import (
	"database/sql"
	"general/constants"
	"general/utils"
	"log"
	"time"
)

func updateHotList() {
	go func() {
		for {
			// 統計所有article的vote和comment數量
			rows, err := connPool.Query(`
			select a.article_id, 
            (select sum(score) from votes v where v.source_id = a.article_id and v.vote_type = 'article')
            as score, count(distinct c.comment_id) as commentCount from articles a 
            left join comments c on c.article_id = a.article_id 
            where a.publish_time <= now() 
            group by a.article_id order by a.publish_time desc;
			`)
			if err != nil {
				log.Println("failed to query hot list!!!")
				time.Sleep(30 * time.Minute)
				continue
			}

			// score * 1 + comment數 * 3
			timeList := []string{}
			hotMap := map[string]int64{}
			for rows.Next() {
				var articleId string
				var sqlScore sql.NullInt64
				var commentCount int16
				err := rows.Scan(&articleId, &sqlScore, &commentCount)
				if err != nil {
					log.Println(err)
					continue
				}
				var score int64
				if sqlScore.Valid {
					val, _ := sqlScore.Value()
					score = val.(int64)
				}
				timeList = append(timeList, articleId)
				num := score + 3*int64(commentCount)
				if num == 0 {
					continue
				}
				hotMap[articleId] = num
			}

			// 照分數排序，把article_id緩存到redis (RPUSH)
			list := utils.SortByIntValue(hotMap)
			uniqueMap := make(map[string]bool)
			for _, id := range list {
				uniqueMap[id] = true
			}

			// 剩下的article照時間順序concat在後
			var filteredList []string
			for _, id := range timeList {
				if _, found := uniqueMap[id]; !found {
					filteredList = append(filteredList, id)
				}
			}
			list = append(list, filteredList...)
			if len(list) == 0 {
				time.Sleep(30 * time.Minute)
				continue
			}

			cache.Del(constants.HOT_LIST_NAMAE)
			err = cache.RPush(constants.HOT_LIST_NAMAE, list)
			if err != nil {
				log.Println("error when pushing hot list name:", err)
			}

			time.Sleep(2 * time.Hour)
		}
	}()
}
