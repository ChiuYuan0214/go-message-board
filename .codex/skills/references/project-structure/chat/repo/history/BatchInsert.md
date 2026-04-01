# HistoryImpl.BatchInsert

**File:** `backend/chat/repo/history.go`
**Struct:** `HistoryImpl`

```go
func (r *HistoryImpl) BatchInsert(chatList []types.DynamoChat) bool
```

- Persists buffered chat messages to DynamoDB
- Used by the scheduler sync flow
- This is the only write path from in-memory chat cache to persistent history
