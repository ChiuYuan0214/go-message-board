# SendMessage

**File:** `backend/chat/services/`

## Signature

```go
func SendMessage(reqMsg *types.RequestEvent)
```

## Behaviour

1. Creates a `Message` from the `RequestEvent`
2. Stores it in the sender's `SendMap` cache (keyed by receiverId)
3. If receiver is online (`IsOnline == true`), writes the message to their WebSocket connection
4. Persistence to DynamoDB happens asynchronously via `SendMap.Sync()`
