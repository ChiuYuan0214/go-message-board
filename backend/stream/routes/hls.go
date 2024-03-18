package routes

import (
	"log"
	"strconv"
	"stream/services"
	"stream/store"
	"stream/utils"

	"github.com/gin-gonic/gin"
	"github.com/notedit/rtmp/format/rtmp"
)

var LiveMap = map[uint64]*rtmp.Conn{}

func HandleHLS(c *gin.Context) {
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
		owner := store.CreateOwner(liveId, userId, conn)
		go services.NewOwnerService(owner).Handle()
	} else {
		store.CreateWatcher(liveId, userId, conn)
	}
}
