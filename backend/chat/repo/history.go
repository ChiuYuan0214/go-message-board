package repo

import (
	"chat/infra"
	"chat/types"
	"time"
)

var _ History = (*HistoryImpl)(nil)

type HistoryImpl struct {
	dynamo infra.Dynamo
}

func (r *HistoryImpl) Run() (err error) {
	return
}

func (r *HistoryImpl) Stop() {}

func (r *HistoryImpl) GetHistory(senderId, receiverId uint64, startTime, endTime time.Time) ([]types.DynamoChat, error) {
	return r.dynamo.Client().GetAllWithFilters(senderId, receiverId, startTime, endTime)
}

func (r *HistoryImpl) GetHistoryLimit20(senderId, receiverId uint64, endTime time.Time) ([]types.DynamoChat, error) {
	return r.dynamo.Client().GetAllWithFilters(senderId, receiverId, time.Now().Add(5*time.Hour), endTime)
}

func (r *HistoryImpl) BatchInsert(chatList []types.DynamoChat) bool {
	return r.dynamo.Client().BatchInsert(chatList)
}
