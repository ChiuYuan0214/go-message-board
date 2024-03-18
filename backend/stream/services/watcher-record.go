package services

import (
	"encoding/json"
	"log"
	"stream/constants"
	"stream/models"
	"stream/types"

	"github.com/gorilla/websocket"
)

type WatcherRecordService struct {
	watcher *types.Watcher
	LiveRecordService
}

func NewWatcherRecordService(watcher *types.Watcher) *WatcherRecordService {
	wrs := &WatcherRecordService{watcher: watcher}
	wrs.SetLive(watcher.GetLive())
	return wrs
}

func (wrs *WatcherRecordService) Handle() {
	watcher := wrs.watcher
	userId := watcher.GetUserId()
	defer watcher.Close()
	for {
		messageType, message, err := watcher.Read()
		if err != nil || messageType != websocket.TextMessage {
			log.Println(err)
			watcher.Close()
			watcher.GetLive().WatcherExit(userId)
			return
		}
		request := models.TextRequest{}
		err = json.Unmarshal(message, &request)
		if err != nil {
			log.Println(err)
			continue
		}
		switch messageType {
		case constants.CHAT, constants.REACT:
			wrs.SendChatOrReaction(userId, message)
		case constants.VOTE:
			wrs.DoVote([]byte(request.Data))
		case constants.FEEDBACK:
			wrs.DoFeedBack([]byte(request.Data))
		}
	}
}

func (wrs *WatcherRecordService) DoVote(message []byte) {
	var request models.VoteRequest
	err := json.Unmarshal(message, &request)
	if err != nil {
		log.Println(err)
		return
	}
	senderId := wrs.watcher.GetUserId()
	vote := &models.Vote{SenderId: senderId, Scores: request.Scores}
	msg, err := json.Marshal(vote)
	if err != nil {
		log.Println(err)
		return
	}
	data := models.TextRequest{DataType: constants.VOTE, Data: string(msg)}
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	wrs.watcher.GetLive().GetOwner().Write(bytes)
}

func (wrs *WatcherRecordService) DoFeedBack(message []byte) {
	var request models.FeedbackRequest
	err := json.Unmarshal(message, &request)
	if err != nil {
		log.Println(err)
		return
	}
	senderId := wrs.watcher.GetUserId()
	feedback := &models.Feedback{SenderId: senderId, Content: request.Content}
	msg, err := json.Marshal(feedback)
	if err != nil {
		log.Println(err)
		return
	}
	data := models.TextRequest{DataType: constants.VOTE, Data: string(msg)}
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	owner := wrs.watcher.GetLive().GetOwner()
	owner.Write(bytes)
}
