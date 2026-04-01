# WatcherService.Handle

**File:** `backend/stream/services/`

## Signature

```go
func (ws *WatcherService) Handle()
```

## Behaviour

Read loop on the watcher's `liveConn`. Watchers are passive — they only receive stream data, never send.
Any message from the watcher closes the connection.
