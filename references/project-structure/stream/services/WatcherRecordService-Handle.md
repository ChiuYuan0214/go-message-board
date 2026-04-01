# WatcherRecordService.Handle

**File:** `backend/stream/services/`

## Signature

```go
func (wrs *WatcherRecordService) Handle()
```

## Behaviour

Read loop on watcher's `recordConn`. Routes by message type:

| Type | Action |
|------|--------|
| `CHAT` / `REACT` | Broadcast to all other watchers |
| `VOTE` | Send poll response to owner (`DoVote`) |
| `FEEDBACK` | Send feedback message to owner (`DoFeedBack`) |
