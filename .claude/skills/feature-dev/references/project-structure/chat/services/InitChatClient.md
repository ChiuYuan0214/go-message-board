# InitChatClient

**File:** `backend/chat/services/`

## Signature

```go
func InitChatClient(conn *websocket.Conn, userId uint64, token string)
```

## Behaviour

1. Creates or updates the client in the global client store
2. Concurrently initialises follower list and follow list (goroutines)
3. Notifies followers/follows that this user is online
4. Stores the token for periodic re-validation

Called once per WebSocket connection, after auth succeeds.
