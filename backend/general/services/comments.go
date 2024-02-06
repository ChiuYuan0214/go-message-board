package services

import (
	"general/types"
	"log"
)

func GetComments(articleId int64) *[]types.CommentListData {
	data := []types.CommentListData{}
	rows, err := connPool.Query(`
	select comment_id, user_id, title, content, creation_time, 
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = 1), 0) as voteUp,
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = -1), 0) as voteDown 
    from comments c where article_id = ? order by creation_time desc;
	`, articleId)
	if err != nil {
		log.Println("failed to get comment list:", err)
		return &data
	}

	for rows.Next() {
		c := types.CommentListData{}
		err = rows.Scan(&c.CommentId, &c.UserId, &c.Title, &c.Content, &c.CreationTime, &c.VoteUp, &c.VoteDown)
		if err != nil {
			log.Println(err)
			continue
		}
		data = append(data, c)
	}

	return &data
}
