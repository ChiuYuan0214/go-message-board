package services

import (
	"chat/repo"
	"chat/store"
	"chat/types"
	"log"
	"strconv"
	"time"
)

var _ History = (*HistoryImpl)(nil)

type HistoryImpl struct {
	historyRepo repo.History
	chatStore   *store.ChatStore
}

func (s *HistoryImpl) Run() (err error) {
	s.chatStore = store.GetChatStore()
	return
}

func (s *HistoryImpl) Stop() {}

func (s *HistoryImpl) GetHistory(event *types.RequestEvent) {
	history := types.History{
		Event: "history",
	}

	timeInt, err := strconv.ParseInt(event.Content, 10, 64)
	if err != nil {
		log.Println(err)
	}

	endTime := time.Now()
	if timeInt != 0 {
		endTime = time.Unix(0, timeInt*1000000)
	}
	startTime := endTime.Add(-5 * time.Hour)

	userHisChan := make(chan *[]types.Message)
	targetHisChan := make(chan *[]types.Message)

	go s.getList(startTime, endTime, event.UserId, event.TargetUserId, userHisChan)
	go s.getList(startTime, endTime, event.TargetUserId, event.UserId, targetHisChan)

	history.TargetUserId = event.TargetUserId
	history.UserHistory = *(<-userHisChan)
	history.TargetHistory = *(<-targetHisChan)

	client, ok := s.chatStore.FindClient(event.UserId)
	if !ok {
		return
	}
	client.Write(history)
}

func (s *HistoryImpl) getList(startTime time.Time, endTime time.Time, senderId uint64, receiverId uint64, channel chan *[]types.Message) {
	sendMap := s.chatStore.GetSendMap(senderId)
	sendMap.Lock.Lock()
	list, cacheStartTime := sendMap.GetCacheMessages(receiverId, startTime, endTime)
	sendMap.Lock.Unlock()

	if len(list) < 10 {
		endTimeForDB := endTime
		if cacheStartTime.Before(endTime) {
			endTimeForDB = cacheStartTime
		}
		chats := s.fetchHistory(senderId, receiverId, startTime, endTimeForDB)
		if len(chats) < 10 {
			chats = s.fetchHistoryLimit20(senderId, receiverId, endTimeForDB)
		}
		dbList := s.translateMessages(chats)
		newList := append(sendMap.GetMessages(receiverId), dbList...)
		filteredList := make([]types.Message, 0)
		count := 0
		for _, m := range newList {
			msgTime := time.Unix(0, m.Time)
			if msgTime.Before(endTime) && (msgTime.After(startTime) || count < 10) {
				filteredList = append(filteredList, m)
				count++
			}
		}
		list = filteredList
		sendMap.Store.Store(receiverId, newList)
	}
	sendMap.MapRef = 0

	channel <- &list
	close(channel)
}

func (s *HistoryImpl) fetchHistory(senderId uint64, receiverId uint64, startTime time.Time, endTime time.Time) []types.DynamoChat {
	chats, err := s.historyRepo.GetHistory(senderId, receiverId, startTime, endTime)
	if err != nil {
		log.Println(err)
		var emptyChats []types.DynamoChat
		return emptyChats
	}
	return chats
}

func (s *HistoryImpl) fetchHistoryLimit20(senderId uint64, receiverId uint64, endTime time.Time) []types.DynamoChat {
	chats, err := s.historyRepo.GetHistoryLimit20(senderId, receiverId, endTime)
	if err != nil {
		log.Println(err)
		var emptyChats []types.DynamoChat
		return emptyChats
	}
	return chats
}

func (s *HistoryImpl) translateMessages(chats []types.DynamoChat) []types.Message {
	var dbList []types.Message
	for _, chat := range chats {
		msg := types.Message{
			UserId:       chat.SenderId,
			TargetUserId: chat.ReceiverId,
			Content:      chat.Content,
			Time:         chat.Time.UnixNano(),
			HasSync:      true,
			Ref:          0,
		}
		dbList = append(dbList, msg)
	}
	return dbList
}
