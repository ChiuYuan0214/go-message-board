# SendMap

**Category:** runtime cache type
**File:** `backend/chat/types/chat.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Lock` | `sync.Mutex` | Guards compound cache operations |
| `Store` | `sync.Map` | `receiverId -> []Message` |
| `MapRef` | `uint8` | Small internal flag used by history sync logic |

## Methods

| Method | Notes |
|--------|-------|
| `Sync(f)` | Runs a function under the map lock |
| `GetMessages(receiverId)` | Returns all cached messages for one target, lazily initializing the bucket |
| `GetCacheMessages(receiverId, startTime, endTime)` | Returns the in-memory history slice for the requested window and the oldest cache timestamp to use for DB fallback |

## Used By

- Each `types.Client` owns one `SendMap`
- `services.MessageImpl` appends new unsynced messages
- `services.HistoryImpl` merges cache with DynamoDB history
- `jobs.SchedulerImpl` uses the cached messages when syncing to DynamoDB
