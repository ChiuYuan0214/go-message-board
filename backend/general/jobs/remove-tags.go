package jobs

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func removeTagsJob() {
	go func() {
		for {
			time.Sleep(time.Hour)

			var tagIds []string
			err := db.Table("tags").Where("tag_id NOT IN (SELECT tag_id FROM article_tag_maps)").Pluck("tag_id", &tagIds).Error
			if err != nil {
				log.Println(err)
				continue
			}
			if len(tagIds) == 0 {
				continue
			}

			stmt := fmt.Sprintf("DELETE FROM tags WHERE tag_id IN (%s)", strings.Join(tagIds, ", "))
			if err = db.Exec(stmt).Error; err != nil {
				log.Println(err)
			}
		}
	}()
}
