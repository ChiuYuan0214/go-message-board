package services

import (
	"chat/types"
	"log"
	"time"
)

func SendMessage(reqMsg *types.RequestEvent) {
	resMsg := types.Message{Event: "message", UserId: reqMsg.UserId, TargetUserId: reqMsg.TargetUserId,
		Content: reqMsg.Content, Time: time.Now().UnixNano(), Ref: 0, HasSync: false}

	from, fromExist := chatStore.GetClient(reqMsg.UserId)
	if !fromExist {
		log.Printf("sender %d not exist.", reqMsg.UserId)
		return
	}
	toward, towardExist := chatStore.GetClient(reqMsg.TargetUserId)
	(*from.SendMap).Lock.Lock()
	sendList, ok := (*from.SendMap).Store[reqMsg.TargetUserId]
	newMsgList := []types.Message{resMsg}
	if !ok {
		(*from.SendMap).Store[reqMsg.TargetUserId] = newMsgList
	} else {
		(*from.SendMap).Store[reqMsg.TargetUserId] = append(newMsgList, sendList...)
	}
	(*from.SendMap).Lock.Unlock()

	if towardExist && toward.IsOnline {
		if !toward.Write(resMsg) {
			log.Printf("failed to send message from %d to %d", from.UserId, toward.UserId)
		}
	}
}
