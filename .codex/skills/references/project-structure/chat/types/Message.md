# Message

**Category:** transport + runtime cache type
**File:** `backend/chat/types/message.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Event` | `string` | Event type identifier, usually `"message"` |
| `UserId` | `uint64` | Sender |
| `TargetUserId` | `uint64` | Receiver |
| `Content` | `string` | Message body |
| `Time` | `int64` | Unix nanoseconds |
| `HasSync` | `bool` | Whether the message already exists in persistent history |
| `Ref` | `uint8` | Small in-memory marker used by cache logic |

## Used By

- Produced by `services.MessageImpl`
- Stored in `types.SendMap`
- Embedded in `types.History`
