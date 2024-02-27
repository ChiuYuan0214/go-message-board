package services

import (
	"database/sql"
	"general/types"
	"log"
)

func GetComments(articleId int64) *[]types.CommentListData {
	data := []types.CommentListData{}
	rows, err := connPool.Query(`
	select c.comment_id, c.user_id, u.username, i.file_name, c.title, c.content, c.creation_time, 
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = 1), 0) as voteUp,
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = -1), 0) as voteDown 
    from comments c inner join users u on u.user_id = c.user_id left join images i on i.user_id = u.user_id 
	where c.article_id = ? order by c.creation_time desc;
	`, articleId)
	if err != nil {
		log.Println("failed to get comment list:", err)
		return &data
	}

	for rows.Next() {
		c := types.CommentListData{}
		var commenterImage sql.NullString
		err = rows.Scan(&c.CommentId, &c.UserId, &c.Commenter, &commenterImage,
			&c.Title, &c.Content, &c.CreationTime, &c.VoteUp, &c.VoteDown)
		if err != nil {
			log.Println(err)
			continue
		}
		if commenterImage.Valid {
			c.CommenterImage = commenterImage.String
		}

		data = append(data, c)
	}

	return &data
}
