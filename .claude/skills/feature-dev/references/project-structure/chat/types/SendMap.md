# SendMap

**Purpose:** Per-client in-memory message cache (before sync to DynamoDB)
**File:** `backend/chat/types/`

## Structure

`sync.Map` keyed by `receiverId` (uint64) → `[]Message`

## Methods

| Method | Notes |
|--------|-------|
| `Sync(receiverId)` | Flush cached messages for target to DynamoDB |
| `GetMessages(receiverId)` | Get all cached messages for target |
| `GetCacheMessages(receiverId)` | Get unsync'd messages only |

## Notes

Each `Client` owns one `SendMap`. Messages are buffered here before being persisted. On history fetch, cache is merged with DynamoDB results.
