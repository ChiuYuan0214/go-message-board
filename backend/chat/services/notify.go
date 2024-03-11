package services

import (
	"chat/types"
)

func NotifyLogin(userId uint64) {
	client, exist := chatStore.GetClient(userId)
	if !exist {
		return
	}

	for _, id := range client.FollowerList {
		follower, ok := chatStore.GetClient(id)
		if !ok || !follower.IsOnline {
			continue
		}
		go follower.Write(types.Notification{Event: "follow-login", UserId: userId})
	}
	for _, id := range client.FollowList {
		follow, ok := chatStore.GetClient(id)
		if !ok || !follow.IsOnline {
			continue
		}
		go follow.Write(types.Notification{Event: "follower-login", UserId: userId})
	}
}

func NotifyLogout(userId uint64) {
	client, exist := chatStore.GetClient(userId)
	if !exist {
		return
	}
	for _, id := range client.FollowerList {
		follower, ok := chatStore.GetClient(id)
		if !ok || !follower.IsOnline {
			continue
		}
		go follower.Write(types.Notification{Event: "follow-logout", UserId: userId})
	}
	for _, id := range client.FollowList {
		follow, ok := chatStore.GetClient(id)
		if !ok || !follow.IsOnline {
			continue
		}
		go follow.Write(types.Notification{Event: "follower-logout", UserId: userId})
	}
}
