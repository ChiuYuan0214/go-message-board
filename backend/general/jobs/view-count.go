package jobs

import (
	"fmt"
	"general/constants"
	"log"
	"time"
)

// run once when boot up
func pushViewCounts() {
	go func() {
		for {
			// terminate job if cache exist
			viewCountMap := cache.HGetAll(constants.VIEW_COUNT_CACHE_NAME)
			if len(viewCountMap) != 0 {
				break
			}
			var results []struct {
				articleId string
				viewCount string
			}
			err := db.Table("articles").Select("article_id, view_count").Find(&results).Error
			if err != nil {
				log.Println("cannot find view_count from sql db:", err)
				time.Sleep(10 * time.Minute)
				continue
			}
			viewCountMap = map[string]string{}
			for _, e := range results {
				viewCountMap[e.articleId] = e.viewCount
			}
			err = cache.HMSet(constants.VIEW_COUNT_CACHE_NAME, &viewCountMap)
			if err != nil {
				log.Println("cannot sync view count to cache!:", err)
			}

			break
		}
	}()
}

// write back view counts toward sql database
func pullViewCounts() {
	go func() {
		for {
			time.Sleep(2 * time.Hour)
			viewCountMap := cache.HGetAll(constants.VIEW_COUNT_CACHE_NAME)
			for id, count := range viewCountMap {
				err := db.Table("articles").Where("article_id = ?", id).Update("view_count", count).Error
				if err != nil {
					log.Println(fmt.Sprintf("failed to sync view count of article %s back to database.", id))
				}
			}
		}
	}()
}
