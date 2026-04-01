# OwnerService.Handle

**File:** `backend/stream/services/`

## Signature

```go
func (ownerService *OwnerService) Handle()
```

## Behaviour

Read loop on the owner's `liveConn` (binary WebSocket). For each message:
- If message type is `LIVE` → calls `LiveService.PushStream` to broadcast to all watchers
- Otherwise → drops the frame

Runs until the owner's connection closes.
