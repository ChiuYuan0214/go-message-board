package services

import (
	"general/types"
	"log"
)

func GetComments(articleId uint64) []types.CommentListData {
	data := []types.CommentListData{}
	err := db.Raw(`select c.comment_id, c.user_id, u.username as commenter, i.file_name as commenterImage, 
	c.title, c.content, c.creation_time as creationTime, 
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = 1), 0) as voteUp,
    coalesce((select count(vote_id) from votes 
    where source_id = c.comment_id and vote_type = 'comment' and score = -1), 0) as voteDown 
    from comments c inner join users u on u.user_id = c.user_id left join images i on i.user_id = u.user_id 
	where c.article_id = ? order by c.creation_time desc`, articleId).Find(&data).Error
	if err != nil {
		log.Println("failed to get comment list:", err)
		return []types.CommentListData{}
	}

	return data
}
