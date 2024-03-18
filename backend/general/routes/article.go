package routes

import (
	"encoding/json"
	"general/routes/middleware"
	"general/services"
	"general/types"
	"general/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func initArticle(router *gin.Engine) {
	ah := ArticleHandler{}
	router.GET("/article", ah.get)
	router.POST("/article", middleware.Auth(), ah.add)
	router.PUT("/article", middleware.Auth(), ah.update)
	router.DELETE("/article", middleware.Auth(), ah.delete)
}

type ArticleHandler struct{}

func (ah *ArticleHandler) get(c *gin.Context) {
	userId := getUserIdFromQuery(c.Request)
	query := getQuery(c.Request)
	id := query.Get("articleId")
	article, status := services.GetArticle(userId, id)
	if status != 0 {
		var message string
		switch status {
		case http.StatusInternalServerError:
			message = "failed to execute query."
		case http.StatusBadRequest:
			message = "id does not exist."
		default:
			message = "something went wrong."
		}
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}

	tagList := services.GetTagsByArticleId(id)
	article.Tags = tagList

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": article})
}

func (ah *ArticleHandler) add(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	var data types.AddArticleData
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if isErr(err) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
		return
	}

	publishTime, err := time.Parse("2006-01-02T15:04", data.PublishTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "time format was wrong."})
		return
	}

	id := services.InsertArticle(userId, &data, &publishTime)
	if id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "failed to insert article."})
		return
	}

	if !services.InsertTags(id, data.Tags) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "failed to insert tags."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "id": id})
}

func (ah *ArticleHandler) update(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	articleId := utils.StringToUint64(getParam(c.Request, "articleId"))
	if articleId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "articleId format incorrect."})
		return
	}

	var data types.UpdateArticleData
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if isErr(err) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
		return
	}

	message, status := services.UpdateArticle(userId, articleId, &data)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}
	if !services.DeleteRemovedTags(articleId, data.Tags) || !services.InsertTags(articleId, data.Tags) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "failed to update tags."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ah *ArticleHandler) delete(c *gin.Context) {
	val, _ := c.Get("userId")
	userId := val.(uint64)
	id := getParam(c.Request, "articleId")

	articleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "articleId format incorrect."})
		return
	}

	message, status := services.DeleteArticle(userId, articleId)
	if message != "" {
		c.JSON(status, gin.H{"status": "fail", "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
