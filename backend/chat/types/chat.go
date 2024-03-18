package types

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type SendMap struct {
	Lock   sync.Mutex
	Store  sync.Map
	MapRef uint8
}

func (sm *SendMap) Sync(f func()) {
	sm.Lock.Lock()
	f()
	sm.Lock.Unlock()
}

func (sm *SendMap) GetMessages(receiverId uint64) []Message {
	_, messageExist := sm.Store.Load(receiverId)
	if !messageExist {
		sm.Store.Store(receiverId, []Message{})
	}
	msgs, _ := sm.Store.Load(receiverId)

	return msgs.([]Message)
}

func (sm *SendMap) GetCacheMessages(receiverId uint64, startTime time.Time, endTime time.Time) ([]Message, time.Time) {
	list := []Message{}
	cache := sm.GetMessages(receiverId)
	for _, msg := range cache {
		msgTime := time.Unix(0, msg.Time)
		if msgTime.After(endTime) || msgTime.Equal(endTime) {
			continue
		}
		if msgTime.Before(startTime) {
			break
		}
		list = append(list, msg)
		msg.Ref = 0
	}

	cacheStartTime := time.Now().Add(1 * time.Minute) // cache最久以前的訊息時間
	cacheSize := len(sm.GetMessages(receiverId))
	cacheIsEmpty := cacheSize == 0

	// 1. cache滿足時間區間 -> cacheStartTime = last cache time / 不用抓dynamo
	// 2. cache部分滿足時間區間 -> cacheStartTime = last cache time / 需要抓dynamo
	// 3. cache存在，但完全不滿足 -> cacheStartTime = last cache time / 需要抓dynamo
	// 4. cache空的 -> cacheStartTime = now / 需要抓dynamo
	if len(list) > 0 {
		cacheStartTime = time.Unix(0, list[len(list)-1].Time)
	} else if !cacheIsEmpty {
		cacheStartTime = time.Unix(0, sm.GetMessages(receiverId)[cacheSize-1].Time) // 如果cache list length > 0 則回傳
	}
	return list, cacheStartTime
}

type Client struct {
	UserId       uint64
	Username     string
	Conn         *websocket.Conn
	ConnLock     sync.Mutex
	SendMap      *SendMap
	FollowerList []uint64
	FollowList   []uint64
	Token        string
	IsOnline     bool
	LogoutTime   time.Time
}

func (c *Client) Logout() {
	c.IsOnline = false
	c.LogoutTime = time.Now()
	c.Conn.Close()
}

func (c *Client) Write(v interface{}) bool {
	c.ConnLock.Lock()
	err := c.Conn.WriteJSON(v)
	c.ConnLock.Unlock()
	if err != nil {
		log.Println(err)
	}
	return err == nil
}

type Chat struct {
	SenderId   uint64    `bson:"senderId"`
	ReceiverId uint64    `bson:"receiverId"`
	Content    string    `bson:"content"`
	Time       time.Time `bson:"time"`
}
