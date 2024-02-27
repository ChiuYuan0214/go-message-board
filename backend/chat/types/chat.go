package types

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type SendMap struct {
	Lock  sync.Mutex
	Store map[uint64][]Message
}

func (sm *SendMap) Sync(f func()) {
	sm.Lock.Lock()
	f()
	sm.Lock.Unlock()
}

func (sm *SendMap) GetMessages(receiverId uint64) []Message {
	_, messageExist := sm.Store[receiverId]
	if !messageExist {
		sm.Store[receiverId] = []Message{}
	}
	return sm.Store[receiverId]
}

func (sm *SendMap) GetCacheMessages(receiverId uint64, startTime time.Time, endTime time.Time) ([]Message, time.Time) {
	list := []Message{}
	for _, msg := range sm.GetMessages(receiverId) {
		msgTime := time.Unix(0, msg.Time)
		if msgTime.After(endTime) || msgTime.Equal(endTime) || msgTime.Before(startTime) {
			break
		}
		list = append(list, msg)
		msg.Ref = 0
	}

	cacheStartTime := time.Now().Add(1 * time.Minute) // cache最久以前的訊息時間
	if len(list) > 0 {
		cacheStartTime = time.Unix(0, list[len(list)-1].Time) // 如果cache list > 0 則回傳
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
