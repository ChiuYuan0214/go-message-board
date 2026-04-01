# RequestEvent

**Purpose:** Client → server WebSocket request
**File:** `backend/chat/types/`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `UserId` | `uint64` | Sender |
| `TargetUserId` | `uint64` | Target (for send/history) |
| `Type` | `string` | Event type: `"send"`, `"history"`, `"addFollow"`, etc. |
| `Content` | `string` | Message content |
