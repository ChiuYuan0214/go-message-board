# NotifyLogout

**File:** `backend/chat/services/`

## Signature

```go
func NotifyLogout(userId uint64)
```

## Behaviour

Broadcasts a `notification` event (`{event: "offline", userId}`) to all online followers and follows.
Called on WebSocket disconnect.
