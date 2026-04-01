# HistoryImpl.GetHistoryLimit20

**File:** `backend/chat/repo/history.go`
**Struct:** `HistoryImpl`

```go
func (r *HistoryImpl) GetHistoryLimit20(senderId, receiverId uint64, endTime time.Time) ([]types.DynamoChat, error)
```

- Uses the same DynamoDB backend as `GetHistory`
- Applies the "fallback to 20 messages" query shape used by the history service
- Used when the requested window plus cache does not produce enough messages
