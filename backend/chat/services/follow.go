package services

import (
	"chat/types"
)

func AddFollow(event *types.RequestEvent) {
	client, ok := chatStore.GetClient(event.UserId)
	if !ok {
		return
	}
	followList := (*client).FollowList
	for _, id := range followList {
		if id == event.TargetUserId {
			return
		}
	}
	(*client).FollowList = append(followList, event.TargetUserId)
	target, ok := chatStore.GetClient(event.TargetUserId)
	if !ok || !target.IsOnline {
		return
	}
	client.Write(types.Notification{Event: "follow-login", UserId: target.UserId})
}

func RemoveFollow(event *types.RequestEvent) {
	client, ok := chatStore.GetClient(event.UserId)
	if !ok {
		return
	}
	followList := (*client).FollowList
	newList := []uint64{}
	for _, id := range followList {
		if id != event.TargetUserId {
			newList = append(newList, id)
		}
	}
	(*client).FollowList = newList
}

func RemoveFollower(event *types.RequestEvent) {
	client, ok := chatStore.GetClient(event.UserId)
	if !ok {
		return
	}
	followerList := (*client).FollowerList
	newList := []uint64{}
	for _, id := range followerList {
		if id != event.TargetUserId {
			newList = append(newList, id)
		}
	}
	(*client).FollowerList = newList
}
