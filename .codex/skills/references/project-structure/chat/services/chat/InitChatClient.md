# ChatImpl.InitChatClient

**File:** `backend/chat/services/chat.go`
**Struct:** `ChatImpl`

```go
func (s *ChatImpl) InitChatClient(conn *websocket.Conn, userId uint64, token string)
```

- Creates or refreshes the in-memory WebSocket client session
- Stores the current connection and token
- Kicks off follow/follower list initialization
- Triggers login notification after setup completes
- Open this first when changing connect-time bootstrap behavior
