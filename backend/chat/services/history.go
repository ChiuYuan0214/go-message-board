package services

import (
	"chat/types"
	"log"
	"strconv"
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
		endTime = time.Unix(0, timeInt*1000000)
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
	sendMap := chatStore.GetSendMap(senderId)
	sendMap.Lock.Lock()
	list, cacheStartTime := sendMap.GetCacheMessages(receiverId, startTime, endTime)
	sendMap.Lock.Unlock()

	if len(list) < 10 { // 如果最舊的cache比指定的時間還晚（靠近現在），則從mongo抓 (cacheStartTime = now() | cache最晚的訊息的時間)
		endTimeForDB := endTime
		if cacheStartTime.Before(endTime) {
			endTimeForDB = cacheStartTime
		}
		chats := fetchHistory(senderId, receiverId, startTime, endTimeForDB)
		if len(chats) < 10 { // 如果資料不足則抓到最多20筆
			chats = fetchHistoryLimit20(senderId, receiverId, endTimeForDB)
		}
		dbList := translateMessages(chats) // 將mongo格式轉換成front-end格式
		newList := append(sendMap.GetMessages(receiverId), dbList...)
		list = []types.Message{}
		count := 0
		for _, m := range newList {
			msgTime := time.Unix(0, m.Time)
			if msgTime.Before(endTime) && (msgTime.After(startTime) || count < 10) {
				list = append(list, m)
				count++
			}
		}
		sendMap.Store.Store(receiverId, newList) // 將mongo的查詢結果加入cache
	}
	sendMap.MapRef = 0

	channel <- &list
	close(channel)
}

func fetchHistory(senderId uint64, receiverId uint64, endTime time.Time, startTime time.Time) []types.Chat {
	condition := bson.D{
		{Key: "senderId", Value: senderId},
		{Key: "receiverId", Value: receiverId},
		{Key: "time", Value: bson.M{"$lt": endTime, "$gte": startTime}},
	}
	opts := options.Find().SetSort(bson.D{{Key: "time", Value: -1}})
	return mongo.FindAll(condition, opts)
}

func fetchHistoryLimit20(senderId uint64, receiverId uint64, endTime time.Time) []types.Chat {
	condition := bson.D{
		{Key: "senderId", Value: senderId},
		{Key: "receiverId", Value: receiverId},
		{Key: "time", Value: bson.M{"$lt": endTime}},
	}
	opts := options.Find().SetSort(bson.D{{Key: "time", Value: -1}}).SetLimit(20)
	chats := mongo.FindAll(condition, opts)
	return chats
}

func translateMessages(chats []types.Chat) []types.Message {
	dbList := []types.Message{}
	for _, chat := range chats {
		msg := types.Message{UserId: chat.SenderId, TargetUserId: chat.ReceiverId,
			Content: chat.Content, Time: chat.Time.UnixNano(), HasSync: true, Ref: 0}
		dbList = append(dbList, msg)
	}
	return dbList
}
