package jobs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func removeTagsJob() {
	go func() {
		for {
			time.Sleep(time.Hour)

			rows, err := connPool.Query("select tag_id from tags where tag_id not exists (select tag_id from article_tag_maps)")
			if err != nil {
				break
			}
			defer rows.Close()

			tagIds := []string{}
			for rows.Next() {
				var tagId int64
				err := rows.Scan(&tagId)
				if err != nil {
					continue
				}
				tagIds = append(tagIds, strconv.FormatInt(tagId, 10))
			}
			if len(tagIds) == 0 {
				continue
			}

			stmt := fmt.Sprintf("delete from tags where tag_id in (%s)", strings.Join(tagIds, ", "))
			connPool.Exec(stmt)
		}
	}()
}
