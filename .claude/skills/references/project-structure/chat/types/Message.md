# Message

**Purpose:** Chat message payload (sent and received)
**File:** `backend/chat/types/`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Event` | `string` | Event type identifier |
| `UserId` | `uint64` | Sender |
| `TargetUserId` | `uint64` | Receiver |
| `Content` | `string` | Message body |
| `Time` | `time.Time` | Send timestamp |
| `HasSync` | `bool` | Whether synced to DB |
| `Ref` | `string` | Reference ID for dedup |
