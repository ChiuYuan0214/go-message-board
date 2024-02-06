package jobs

import (
	"fmt"
	"general/constants"
	"log"
	"strconv"
	"strings"
	"time"
)

func updateTopComments() {
	go func() {
		for {
			// 根據set查找有被comment過的article
			list := cache.SMembers(constants.COMMENTED_ARTICLE_SET)
			// 消除redis的紀錄
			err := cache.Del(constants.COMMENTED_ARTICLE_SET)
			if err != nil {
				log.Println("failed to remove hot list cache!")
			}
			if len(list) == 0 {
				time.Sleep(time.Hour)
				continue
			}

			// 把這些article的對應comment_id, vote_score查出來
			stmt := fmt.Sprintf(`
	        select c.comment_id, c.article_id, sum(v.score) from votes v 
            inner join comments c on c.comment_id = v.source_id and v.vote_type = 'comment'
			where c.article_id in (%s) 
            group by c.comment_id`, strings.Join(list, ", "))
			rows, err := connPool.Query(stmt)
			if err != nil {
				log.Println("failed to query comment scores!:", err)
			}

			// 挑選score總分最高的comment作為top comment
			commentScoreMap := map[uint64]uint16{}
			articleCommentMap := map[uint64]uint64{}
			for rows.Next() {
				var commentId uint64
				var articleId uint64
				var score uint16
				err := rows.Scan(&commentId, &articleId, &score)
				if err != nil {
					log.Println("failed to query comment scores!:", err)
				}

				prevCommentId, exist := articleCommentMap[articleId]
				if !exist || score > commentScoreMap[prevCommentId] {
					commentScoreMap[commentId] = score
					articleCommentMap[articleId] = commentId
				}
			}

			for key, val := range articleCommentMap {
				_, err := connPool.Exec("update articles set top_comment_id = ? where article_id = ?", val, key)
				if err != nil {
					log.Println(fmt.Sprintf(`
					failed to update top comment to commentId %s of articleId %s`,
						strconv.FormatUint(val, 10), strconv.FormatUint(key, 10)))
				}
			}

			time.Sleep(time.Hour)
		}
	}()
}
