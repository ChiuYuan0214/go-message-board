package services

import (
	"encoding/json"
	"fmt"
	"log"
	"stream/constants"
	"stream/models"
	"stream/types"
	"strings"

	"github.com/gorilla/websocket"
)

type OwnerRecordService struct {
	owner *types.Owner
	LiveRecordService
}

func NewOwnerRecordService(owner *types.Owner) *OwnerRecordService {
	ownerRecordService := &OwnerRecordService{owner: owner}
	ownerRecordService.SetLive(owner.GetLive())
	return ownerRecordService
}

func (ors *OwnerRecordService) Handle() {
	owner := ors.owner
	userId := owner.GetUserId()
	defer ors.owner.Close()
	owner.Write([]byte("connection success"))
	for {
		messageType, message, err := owner.Read()
		if err != nil || messageType != websocket.TextMessage {
			log.Println(err)
			owner.Close()
			owner.GetLive().OwnerExit()
			return
		}
		request := models.TextRequest{}
		err = json.Unmarshal(message, &request)
		if err != nil {
			log.Println(err)
			continue
		}

		switch request.DataType {
		case constants.CHAT, constants.REACT:
			ors.SendChatOrReaction(userId, message)
		case constants.VOTE:
			ors.OpenSubject(message, []byte(request.Data))
		case constants.FEEDBACK:
			ors.ResponseFeedback(message, []byte(request.Data))
		}
	}
}

func (ors *OwnerRecordService) OpenSubject(message, data []byte) {
	var request models.SubjectRequest
	err := json.Unmarshal(data, &request)
	if err != nil {
		log.Println(err)
		return
	} else if strings.Trim(request.Title, " ") == "" ||
		len(request.Questions) == 0 ||
		strings.Trim(request.Questions[0], " ") == "" {
		return
	}
	live := ors.owner.GetLive()
	watchers := live.GetWatchers()
	watchers.Range(func(key, val any) bool {
		id := key.(uint64)
		watcher := val.(*types.Watcher)
		if id == ors.owner.GetUserId() {
			return true
		}
		go func() {
			err := watcher.Write(message)
			if err != nil {
				log.Println(err)
				watcher.Close()
				live.WatcherExit(id)
			}
		}()
		return true
	})
}

func (ors *OwnerRecordService) ResponseFeedback(message, data []byte) {
	var request models.ResponseFeedbackRequest
	err := json.Unmarshal(data, &request)
	if err != nil {
		log.Println(err)
		return
	}

	if !ors.live.HasWatcher(request.ReceiverId) {
		msgMap := map[string]string{"error": fmt.Sprintf("failed to send response toward %d, watcher not exist.", request.ReceiverId)}
		msg, err := json.Marshal(msgMap)
		if err != nil {
			return
		}
		ors.owner.Write(msg)
		return
	}
	val, ok := ors.live.GetWatchers().Load(request.ReceiverId)
	if !ok {
		return
	}
	watcher := val.(*types.Watcher)
	err = watcher.Write(message)
	if err != nil {
		log.Println(err)
	}
}
