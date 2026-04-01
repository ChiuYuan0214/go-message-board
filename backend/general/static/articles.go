package static

var BaseSearchArticlesStmt = `
    select a.article_id, a.user_id, a.title, a.content, u.username, i.file_name, 
	coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = 1 group by source_id), 0) as voteUp, 
    coalesce((select count(vote_id) from votes 
    where source_id = a.article_id and vote_type = 'article' and score = -1 group by source_id), 0) as voteDown, 
	ifnull((select score from votes where source_id = a.article_id and vote_type = 'article' and user_id = ?), 0) as myScore, 
	((select a.article_id from collections collec where collec.article_id = a.article_id and collec.user_id = ?) is not null) as hasCollec, 
	c.title, c.content, cu.username, ci.file_name, a.publish_time 
	from articles a 
	inner join users u on u.user_id = a.user_id 
	left join comments c on top_comment_id = comment_id 
	left join users cu on cu.user_id = c.user_id 
	left join images i on i.user_id = u.user_id 
	left join images ci on ci.user_id = c.user_id 
	where a.publish_time <= now() `
