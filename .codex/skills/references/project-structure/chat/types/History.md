# History

**Category:** transport response type
**File:** `backend/chat/types/history.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Event` | `string` | Always `"history"` |
| `TargetUserId` | `uint64` | The other user in the conversation |
| `UserHistory` | `[]Message` | Messages sent by the caller |
| `TargetHistory` | `[]Message` | Messages sent by the target |

## Used By

- Built by `services.HistoryImpl.GetHistory`
- Written back to the requester over the WebSocket connection
