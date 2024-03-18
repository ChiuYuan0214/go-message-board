package routes

import (
	"log"
	"strconv"
	"stream/services"
	"stream/store"
	"stream/utils"

	"github.com/gin-gonic/gin"
)

func HandleSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	token := utils.GetTokenFromQuery(c.Request)
	userId, err := strconv.ParseUint(token, 10, 64)
	if err != nil {
		return
	}

	liveId := utils.GetLiveIdFromQuery(c.Request)
	isLiveOwner := userId == 123 && liveId == 1

	if isLiveOwner {
		owner := store.CreateOwnerRecord(liveId, userId, conn)
		go services.NewOwnerRecordService(owner).Handle()
	} else {
		watcher := store.CreateWatcherRecord(liveId, userId, conn)
		go services.NewWatcherRecordService(watcher).Handle()
	}
}
