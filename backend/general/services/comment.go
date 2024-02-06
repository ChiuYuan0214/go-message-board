package services

import (
	"fmt"
	"general/constants"
	"general/types"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AddComment(userId int64, data *types.AddCommentData) int64 {
	row := connPool.QueryRow("select count(article_id) from articles where article_id = ?", data.ArticleId)
	var count int
	if err := row.Scan(&count); err != nil || count < 1 {
		return 0
	}
	sqlRes, err := connPool.Exec("insert into comments (user_id, article_id, title, content) values (?, ?, ?, ?)", userId, data.ArticleId, data.Title, data.Content)
	if err != nil {
		log.Println("err of inserting comment:", err)
		return 0
	}
	commentId, _ := sqlRes.LastInsertId()

	recordArticle(data.ArticleId)

	return commentId
}

func UpdateComment(userId int64, data *types.UpdateCommentData) (string, int) {
	row := connPool.QueryRow("select user_id from comments where comment_id = ?", data.CommentId)

	var actualUserId int64
	err := row.Scan(&actualUserId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	if userId != actualUserId {
		return "user incorrect.", http.StatusBadRequest
	}

	isEdited := false
	stmt := "update comments set "
	if strings.Trim(data.Title, " ") != "" {
		stmt += fmt.Sprintf(" title = '%s', ", data.Title)
		isEdited = true
	}

	if strings.Trim(data.Content, " ") != "" {
		stmt += fmt.Sprintf(" content = '%s', ", data.Content)
		isEdited = true
	}

	if !isEdited {
		return "nothing to update.", http.StatusBadRequest
	}

	stmt += "edited = true where comment_id = ?"
	_, err = connPool.Exec(stmt, data.CommentId)
	if err != nil {
		return "something went wrong", http.StatusInternalServerError
	}

	return "", 0
}

func DeleteComment(userId int64, commentId string) (string, int) {
	row := connPool.QueryRow("select user_id from comments where comment_id = ?", commentId)
	var actualUserId int64
	err := row.Scan(&actualUserId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}
	if userId != actualUserId {
		return "user incorrect.", http.StatusBadRequest
	}

	// delete related votes
	_, err = connPool.Exec("delete from votes where vote_type = 'comment' and source_id = ?", commentId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	// delete comment
	_, err = connPool.Exec("delete from comments where comment_id = ?", commentId)
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	return "", 0
}

func recordArticle(articleId int64) error {
	id := strconv.FormatInt(articleId, 10)
	return cache.SAdd(constants.COMMENTED_ARTICLE_SET, id)
}
