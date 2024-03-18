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
			var results []struct {
				commentId uint64
				articleId uint64
				score     uint16
			}
			err = db.Table("votes").Select("comments.comment_id, comments.article_id, SUM(votes.score)").
				Joins("INNER JOIN comments ON comments.comment_id = votes.source_id AND votes.vote_type = 'comment'").
				Where("comments.article_id IN (?)", strings.Join(list, ", ")).
				Group("comments.comment_id").
				Find(&results).Error
			if err != nil {
				log.Println(err)
				time.Sleep(time.Hour)
				continue
			}

			// 挑選score總分最高的comment作為top comment
			commentScoreMap := map[uint64]uint16{}
			articleCommentMap := map[uint64]uint64{}
			for _, e := range results {
				prevCommentId, exist := articleCommentMap[e.articleId]
				if !exist || e.score > commentScoreMap[prevCommentId] {
					commentScoreMap[e.commentId] = e.score
					articleCommentMap[e.articleId] = e.commentId
				}
			}

			for key, val := range articleCommentMap {
				err := db.Table("articles").Where("article_id = ?", key).Update("top_comment_id", val).Error
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
