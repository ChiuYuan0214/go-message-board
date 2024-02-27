package services

import (
	"chat/types"
	"log"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetHistory(event *types.RequestEvent) {
	history := types.History{Event: "history"}
	timeInt, err := strconv.ParseInt(event.Content, 10, 64)
	if err != nil {
		log.Println(err)
	}
	endTime := time.Now()
	if timeInt != 0 {
		endTime = time.Unix(0, timeInt)
	}
	startTime := endTime.Add(-5 * time.Hour)
	userHisChan := make(chan *[]types.Message)
	targetHisChan := make(chan *[]types.Message)
	go getList(startTime, endTime, event.UserId, event.TargetUserId, userHisChan)
	go getList(startTime, endTime, event.TargetUserId, event.UserId, targetHisChan)

	history.TargetUserId = event.TargetUserId
	history.UserHistory = *(<-userHisChan)
	history.TargetHistory = *(<-targetHisChan)

	client, ok := chatStore.GetClient(event.UserId)
	if !ok {
		return
	}
	client.Write(history)
}

func getList(startTime time.Time, endTime time.Time, senderId uint64, receiverId uint64, channel chan *[]types.Message) {
	_, clientExist := chatStore.GetClient(senderId)
	if !clientExist {
		(*chatStore.Clients)[senderId] = &types.Client{UserId: senderId,
			SendMap: &types.SendMap{Lock: sync.Mutex{}, Store: map[uint64][]types.Message{}}}
	}

	sendMap := chatStore.GetSendMap(senderId)
	sendMap.Lock.Lock()
	list, cacheStartTime := sendMap.GetCacheMessages(receiverId, startTime, endTime)
	hasCache := len(list) > 0

	if cacheStartTime.After(startTime) { // 如果最舊的cache比指定的時間還晚（靠近現在），則從mongo抓 (cacheStartTime = now() | cache最晚的訊息的時間)
		if hasCache {
			endTime = cacheStartTime
		}
		chats := fetchHistory(senderId, receiverId, startTime, endTime)
		if len(*chats) < 10 { // 如果資料不足則抓到最多20筆
			chats = fetchHistoryLimit20(senderId, receiverId, endTime)
		}
		dbList := translateMessages(chats)                                        // 將mongo格式轉換成front-end格式
		list = append(list, *dbList...)                                           // 將mongo的查詢結果加入給front-end的result
		sendMap.Store[receiverId] = append(sendMap.Store[receiverId], *dbList...) // 將mongo的查詢結果加入cache
	}
	sendMap.Lock.Unlock()
	channel <- &list
	close(channel)
}

func fetchHistory(senderId uint64, receiverId uint64, endTime time.Time, startTime time.Time) *[]types.Chat {
	condition := bson.D{
		{Key: "senderId", Value: senderId},
		{Key: "receiverId", Value: receiverId},
		{Key: "time", Value: bson.M{"$lt": endTime, "$gte": startTime}},
	}
	opts := options.Find().SetSort(bson.D{{Key: "time", Value: -1}})
	return mongo.FindAll(condition, opts)
}

func fetchHistoryLimit20(senderId uint64, receiverId uint64, endTime time.Time) *[]types.Chat {
	condition := bson.D{
		{Key: "senderId", Value: senderId},
		{Key: "receiverId", Value: receiverId},
		{Key: "time", Value: bson.M{"$lt": endTime}},
	}
	opts := options.Find().SetSort(bson.D{{Key: "time", Value: -1}}).SetLimit(20)
	chats := mongo.FindAll(condition, opts)
	return chats
}

func translateMessages(chats *[]types.Chat) *[]types.Message {
	dbList := []types.Message{}
	for _, chat := range *chats {
		msg := types.Message{UserId: chat.SenderId, TargetUserId: chat.ReceiverId,
			Content: chat.Content, Time: chat.Time.UnixNano(), HasSync: true, Ref: 0}
		dbList = append(dbList, msg)
	}
	return &dbList
}
