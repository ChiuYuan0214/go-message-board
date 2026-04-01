# History

**Purpose:** Server → client history response
**File:** `backend/chat/types/`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Event` | `string` | Always `"history"` |
| `TargetUserId` | `uint64` | The other user in the conversation |
| `UserHistory` | `[]Message` | Messages sent by the caller |
| `TargetHistory` | `[]Message` | Messages sent by the target |
