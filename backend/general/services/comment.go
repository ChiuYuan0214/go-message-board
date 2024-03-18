package services

import (
	"general/constants"
	"general/entities"
	"general/types"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AddComment(userId uint64, data *types.AddCommentData) uint64 {
	var count int64
	err := db.Model(&types.Article{}).Where("article_id = ?", data.ArticleId).Count(&count).Error
	if err != nil || count < 1 {
		return 0
	}
	newComment := entities.Comment{UserId: userId, ArticleId: data.ArticleId, Title: data.Title, Content: data.Content}
	result := db.Create(&newComment)
	if result.Error != nil {
		log.Println("err of inserting comment:", result.Error)
		return 0
	}

	recordArticle(data.ArticleId)

	return newComment.CommentId
}

func UpdateComment(userId uint64, data *types.UpdateCommentData) (string, int) {
	var actualUserId uint64
	err := db.Model(&entities.Comment{}).Where("comment_id = ?", data.CommentId).Scan(&actualUserId).Error
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	if userId != actualUserId {
		return "user incorrect.", http.StatusBadRequest
	}

	newComment := entities.Comment{CommentId: data.CommentId, Title: data.Title, Content: data.Content}
	fields := map[string]string{}
	if strings.Trim(data.Title, " ") != "" {
		fields["title"] = data.Title
	}
	if strings.Trim(data.Content, " ") != "" {
		fields["content"] = data.Content
	}
	if len(fields) == 0 {
		return "nothing to update.", http.StatusBadRequest
	}

	err = db.Model(newComment).Updates(fields).Error
	if err != nil {
		return "something went wrong", http.StatusInternalServerError
	}

	return "", 0
}

func DeleteComment(userId uint64, commentId string) (string, int) {
	var actualUserId uint64
	err := db.Model(&entities.Comment{}).Select("user_id").Where("comment_id = ?", commentId).Scan(&actualUserId).Error
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}
	if userId != actualUserId {
		return "user incorrect.", http.StatusBadRequest
	}

	// delete related votes
	err = db.Where("vote_type = 'comment' and source_id = ?", commentId).Delete(&entities.Vote{}).Error
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	// delete comment
	err = db.Where("comment_id = ?", commentId).Delete(&entities.Comment{}).Error
	if err != nil {
		return "something went wrong.", http.StatusInternalServerError
	}

	return "", 0
}

func recordArticle(articleId uint64) error {
	id := strconv.FormatUint(articleId, 10)
	return cache.SAdd(constants.COMMENTED_ARTICLE_SET, id)
}
