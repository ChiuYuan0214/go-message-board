# Client

**Purpose:** Represents a connected WebSocket user
**File:** `backend/chat/types/`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `UserId` | `uint64` | |
| `Username` | `string` | |
| `Conn` | `*websocket.Conn` | WebSocket connection |
| `ConnLock` | `sync.Mutex` | Guards concurrent writes to Conn |
| `SendMap` | `*SendMap` | Per-target message cache |
| `FollowerList` | `[]uint64` | User IDs of followers |
| `FollowList` | `[]uint64` | User IDs this client follows |
| `Token` | `string` | For periodic re-validation |
| `IsOnline` | `bool` | |
| `LogoutTime` | `time.Time` | Set on disconnect |

## Methods

| Method | Notes |
|--------|-------|
| `Logout()` | Marks offline, sets LogoutTime |
| `Write(msg)` | Thread-safe write via ConnLock |

## Notes

Stored in a global client map keyed by userId. Checking `IsOnline` tells you if a target user can receive messages.
