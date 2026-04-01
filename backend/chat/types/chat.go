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
		var messages []Message
		sm.Store.Store(receiverId, messages)
	}
	msgs, _ := sm.Store.Load(receiverId)

	return msgs.([]Message)
}

func (sm *SendMap) GetCacheMessages(receiverId uint64, startTime time.Time, endTime time.Time) ([]Message, time.Time) {
	var list []Message
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
	StateLock    sync.RWMutex
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

func (c *Client) InitSession(conn *websocket.Conn, token string) {
	c.StateLock.Lock()
	c.Conn = conn
	c.Token = token
	c.IsOnline = true
	c.StateLock.Unlock()
}

func (c *Client) ReplaceFollowerList(list []uint64) {
	c.StateLock.Lock()
	c.FollowerList = list
	c.StateLock.Unlock()
}

func (c *Client) ReplaceFollowList(list []uint64) {
	c.StateLock.Lock()
	c.FollowList = list
	c.StateLock.Unlock()
}

func (c *Client) SnapshotFollowerList() []uint64 {
	c.StateLock.RLock()
	defer c.StateLock.RUnlock()

	return append([]uint64(nil), c.FollowerList...)
}

func (c *Client) SnapshotFollowList() []uint64 {
	c.StateLock.RLock()
	defer c.StateLock.RUnlock()

	return append([]uint64(nil), c.FollowList...)
}

func (c *Client) AddFollow(targetUserId uint64) bool {
	c.StateLock.Lock()
	defer c.StateLock.Unlock()

	for _, id := range c.FollowList {
		if id == targetUserId {
			return false
		}
	}

	c.FollowList = append(c.FollowList, targetUserId)
	return true
}

func (c *Client) RemoveFollow(targetUserId uint64) {
	c.StateLock.Lock()
	defer c.StateLock.Unlock()

	var newList []uint64
	for _, id := range c.FollowList {
		if id != targetUserId {
			newList = append(newList, id)
		}
	}
	c.FollowList = newList
}

func (c *Client) RemoveFollower(targetUserId uint64) {
	c.StateLock.Lock()
	defer c.StateLock.Unlock()

	var newList []uint64
	for _, id := range c.FollowerList {
		if id != targetUserId {
			newList = append(newList, id)
		}
	}
	c.FollowerList = newList
}

func (c *Client) TokenValue() string {
	c.StateLock.RLock()
	defer c.StateLock.RUnlock()

	return c.Token
}

func (c *Client) IsActive() bool {
	c.StateLock.RLock()
	defer c.StateLock.RUnlock()

	return c.IsOnline
}

func (c *Client) LogoutSnapshot() (bool, time.Time) {
	c.StateLock.RLock()
	defer c.StateLock.RUnlock()

	return c.IsOnline, c.LogoutTime
}

func (c *Client) CloseSession() {
	c.StateLock.Lock()
	c.IsOnline = false
	c.LogoutTime = time.Now()
	conn := c.Conn
	c.StateLock.Unlock()

	if conn != nil {
		conn.Close()
	}
}

func (c *Client) Read(v interface{}) error {
	c.StateLock.RLock()
	conn := c.Conn
	c.StateLock.RUnlock()

	if conn == nil {
		return websocket.ErrCloseSent
	}

	return conn.ReadJSON(v)
}

func (c *Client) Write(v interface{}) bool {
	c.StateLock.RLock()
	conn := c.Conn
	c.StateLock.RUnlock()

	if conn == nil {
		return false
	}

	c.ConnLock.Lock()
	err := conn.WriteJSON(v)
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
