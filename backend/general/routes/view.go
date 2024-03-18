package routes

import (
	"general/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initView(router *gin.Engine) {
	vh := ViewHandler{}
	router.PUT("/view", vh.record)
}

type ViewHandler struct{}

func (vh *ViewHandler) record(c *gin.Context) {
	articleId := getParam(c.Request, "articleId")
	services.RecordView(articleId)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
