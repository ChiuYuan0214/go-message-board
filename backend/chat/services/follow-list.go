package services

import (
	"chat/store"
	"chat/types"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func InitFollowerList(wg *sync.WaitGroup, conn *websocket.Conn, userId uint64) {
	clients := store.GetChatStore().Clients
	rows, err := connPool.Query(`select follower_id from followers where user_id = ?`, userId)
	if err != nil {
		log.Println(err)
	}
	followerList := []uint64{}
	onlineList := []uint64{}
	for rows.Next() {
		var followerId uint64
		err = rows.Scan(&followerId)
		if err != nil {
			log.Println(err)
			continue
		}
		followerList = append(followerList, followerId)
		follower, ok := (*clients)[followerId]
		if !ok || !follower.IsOnline {
			continue
		}
		onlineList = append(onlineList, followerId)
	}
	client := (*clients)[userId]
	client.FollowerList = followerList
	client.Write(types.UserInfoList{Event: "online-follower-list", List: onlineList})
	(*wg).Done()
}

func InitFollowList(wg *sync.WaitGroup, conn *websocket.Conn, userId uint64) {
	clients := store.GetChatStore().Clients
	rows, err := connPool.Query(`select user_id from followers where follower_id = ?`, userId)
	if err != nil {
		log.Println(err)
	}
	followList := []uint64{}
	onlineList := []uint64{}
	for rows.Next() {
		var followId uint64
		err = rows.Scan(&followId)
		if err != nil {
			log.Println(err)
			continue
		}
		followList = append(followList, followId)
		follow, ok := (*clients)[followId]
		if !ok || !follow.IsOnline {
			continue
		}
		onlineList = append(onlineList, followId)
	}
	client := (*clients)[userId]
	client.FollowList = followList

	client.Write(types.UserInfoList{Event: "online-follow-list", List: onlineList})
	(*wg).Done()
}
