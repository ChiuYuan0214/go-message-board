# NotifyLogin

**File:** `backend/chat/services/`

## Signature

```go
func NotifyLogin(userId uint64)
```

## Behaviour

Broadcasts a `notification` event (`{event: "online", userId}`) to all online followers and follows of the user.
Called by [InitChatClient](InitChatClient.md) on connect.
