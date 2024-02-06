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

			rows, err := connPool.Query("select article_id, view_count from articles")
			if err != nil {
				log.Println("cannot sync view count to cache!:", err)
			}
			viewCountMap = map[string]string{}
			for rows.Next() {
				var articleId string
				var viewCount string
				rows.Scan(&articleId, &viewCount)
				viewCountMap[articleId] = viewCount
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
				_, err := connPool.Exec("update articles set view_count = ? where article_id = ?", count, id)
				if err != nil {
					log.Println(fmt.Sprintf("failed to sync view count of article %s back to database.", id))
				}
			}
		}
	}()
}
