package services

import (
	"chat/store"
	"chat/types"
)

var _ Event = (*EventImpl)(nil)

type EventImpl struct {
	messageService Message
	followService  Follow
	historyService History
	chatStore      *store.ChatStore
}

func (s *EventImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	return
}

func (s *EventImpl) Stop() {}

func (s *EventImpl) RunEventLoop() {
	broadcast := s.chatStore.Broadcast
	for {
		event := <-broadcast
		go s.HandleEvent(event)
	}
}

func (s *EventImpl) HandleEvent(event *types.RequestEvent) {
	switch event.Type {
	case "message":
		s.messageService.SendMessage(event)
	case "add-follow":
		s.followService.AddFollow(event)
	case "remove-follow":
		s.followService.RemoveFollow(event)
	case "remove-follower":
		s.followService.RemoveFollower(event)
	case "refresh-token":
		// no-op for now
	case "history":
		s.historyService.GetHistory(event)
	}
}
