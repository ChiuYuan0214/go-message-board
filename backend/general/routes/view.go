package routes

import (
	"general/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewHandler struct {
	router      Router
	viewService service.View
}

func (vh *ViewHandler) Run() (err error) {
	vh.router.Put("/view", vh.record)
	return
}

func (vh *ViewHandler) Stop() {}

func (vh *ViewHandler) record(c *gin.Context) {
	articleId := getParam(c.Request, "articleId")
	vh.viewService.RecordView(articleId)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
