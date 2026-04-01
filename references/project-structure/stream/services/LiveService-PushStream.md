# LiveService.PushStream

**File:** `backend/stream/services/`

## Signature

```go
func (c *LiveService) PushStream(senderId uint64, watchers *types.WatcherMap, message []byte)
```

## Parameters

| Name | Notes |
|------|-------|
| `senderId` | Excluded from broadcast (don't echo to sender) |
| `watchers` | The live session's watcher map |
| `message` | Raw binary frame to broadcast |

## Behaviour

Iterates all watchers in the `WatcherMap` and writes the binary message to each one's `liveConn`, skipping the sender.
