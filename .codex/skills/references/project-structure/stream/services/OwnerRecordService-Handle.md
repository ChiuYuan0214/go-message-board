# OwnerRecordService.Handle

**File:** `backend/stream/services/`

## Signature

```go
func (ors *OwnerRecordService) Handle()
```

## Behaviour

Read loop on owner's `recordConn` (text WebSocket). Routes by message type:

| Type | Action |
|------|--------|
| `CHAT` / `REACT` | Broadcast to all watchers via [SendChatOrReaction](LiveService-PushStream.md) |
| `VOTE` | Sends poll/survey to all watchers via `OpenSubject` |
| `FEEDBACK` | Relays owner's feedback response to a specific watcher |
