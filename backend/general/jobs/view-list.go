package jobs

import (
	"general/constants"
	"general/utils"
	"log"
	"time"
)

// update view list cache from view count cache
func updateViewList() {
	go func() {
		for {
			viewCountMap := cache.HGetAll(constants.VIEW_COUNT_CACHE_NAME)
			if len(viewCountMap) == 0 {
				time.Sleep(20 * time.Second)
				continue
			}

			list := utils.SortByValue(viewCountMap)
			cache.Del(constants.VIEW_LIST_NAME)
			err := cache.RPush(constants.VIEW_LIST_NAME, list)
			if err != nil {
				log.Println(err)
			}

			time.Sleep(2 * time.Hour)
		}
	}()
}
