package jobs

import (
	"general/constants"
	"general/utils"
	"log"
	"time"
)

func updateHotList() {
	go func() {
		for {
			// 統計所有article的vote和comment數量
			var list []struct {
				articleId    string
				score        uint16
				commentCount uint16
			}
			err := db.Raw(`
			select a.article_id, 
            (select sum(score) from votes v where v.source_id = a.article_id and v.vote_type = 'article')
            as score, count(distinct c.comment_id) as commentCount from articles a 
            left join comments c on c.article_id = a.article_id 
            where a.publish_time <= now() 
            group by a.article_id;
			`).Find(&list).Error
			if err != nil {
				log.Println("failed to query hot list!!!")
				time.Sleep(30 * time.Minute)
				continue
			}

			// score * 1 + comment數 * 3
			timeList := []string{}
			hotMap := map[string]uint16{}
			for _, e := range list {
				timeList = append(timeList, e.articleId)
				num := e.score + 3*e.commentCount
				if num == 0 {
					continue
				}
				hotMap[e.articleId] = num
			}

			// 照分數排序，把article_id緩存到redis (RPUSH)
			sortedList := utils.SortByIntValue(hotMap)
			uniqueMap := make(map[string]bool)
			for _, id := range sortedList {
				uniqueMap[id] = true
			}

			// 剩下的article照時間順序concat在後
			var filteredList []string
			for _, id := range timeList {
				if _, found := uniqueMap[id]; !found {
					filteredList = append(filteredList, id)
				}
			}
			sortedList = append(sortedList, filteredList...)
			if len(sortedList) == 0 {
				time.Sleep(30 * time.Minute)
				continue
			}

			cache.Del(constants.HOT_LIST_NAMAE)
			err = cache.RPush(constants.HOT_LIST_NAMAE, sortedList)
			if err != nil {
				log.Println("error when pushing hot list name:", err)
			}

			time.Sleep(2 * time.Hour)
		}
	}()
}
