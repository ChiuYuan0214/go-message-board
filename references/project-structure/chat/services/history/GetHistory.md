# HistoryImpl.GetHistory

**File:** `backend/chat/services/history.go`
**Struct:** `HistoryImpl`

```go
func (s *HistoryImpl) GetHistory(event *types.RequestEvent)
```

- Fetches conversation history for the requested window
- Merges DynamoDB history with in-memory unsynced cache
- Writes a `History` response back to the requesting client
- Open this when changing history pagination, cache merge, or Dynamo fallback behavior
