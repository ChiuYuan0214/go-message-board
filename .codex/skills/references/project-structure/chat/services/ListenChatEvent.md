# ListenChatEvent

**File:** `backend/chat/services/`

## Signature

```go
func ListenChatEvent(ctx context.Context, cancel context.CancelFunc, userId uint64)
```

## Behaviour

Main WebSocket read loop. Reads JSON messages from the client, routes them by `Type`:
- `send` → [SendMessage](SendMessage.md)
- `history` → [GetHistory](GetHistory.md)
- `addFollow` / `removeFollow` / `removeFollower` → follow management
- `ping` → ignored

On disconnect or context cancel, calls [NotifyLogout](NotifyLogout.md) and cleans up the client.
