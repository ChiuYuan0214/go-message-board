package repo

import (
	"chat/types"
	"time"
)

type Token interface {
	GetToken(userId uint64) (string, error)
}

type Follow interface {
	GetFollowerIDs(userId uint64) ([]uint64, error)
	GetFollowIDs(userId uint64) ([]uint64, error)
}

type History interface {
	GetHistory(senderId, receiverId uint64, startTime, endTime time.Time) ([]types.DynamoChat, error)
	GetHistoryLimit20(senderId, receiverId uint64, endTime time.Time) ([]types.DynamoChat, error)
	BatchInsert(chatList []types.DynamoChat) bool
}
