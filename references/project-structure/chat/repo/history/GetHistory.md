# HistoryImpl.GetHistory

**File:** `backend/chat/repo/history.go`
**Struct:** `HistoryImpl`

```go
func (r *HistoryImpl) GetHistory(senderId, receiverId uint64, startTime, endTime time.Time) ([]types.DynamoChat, error)
```

- Reads DynamoDB chat history for a sender/receiver pair within the requested window
- Primary history read path used before the "limit 20" fallback
