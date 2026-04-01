# ChatImpl.ListenChatEvent

**File:** `backend/chat/services/chat.go`
**Struct:** `ChatImpl`

```go
func (s *ChatImpl) ListenChatEvent(ctx context.Context, cancel context.CancelFunc, userId uint64)
```

- Runs the main read loop for one connected client
- Rejects mismatched `msg.UserId`
- Pushes accepted events into the shared broadcast channel
- On disconnect, logs out the client and triggers logout notifications
- Owns the per-connection read loop, not event dispatch
