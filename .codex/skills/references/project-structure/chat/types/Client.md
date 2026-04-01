# Client

**Category:** runtime connection state
**File:** `backend/chat/types/chat.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `UserId` | `uint64` | Connected user ID |
| `Username` | `string` | Present but not actively populated in current flow |
| `Conn` | `*websocket.Conn` | WebSocket connection |
| `ConnLock` | `sync.Mutex` | Guards concurrent writes to Conn |
| `SendMap` | `*SendMap` | Per-target message cache |
| `FollowerList` | `[]uint64` | User IDs of followers |
| `FollowList` | `[]uint64` | User IDs this client follows |
| `Token` | `string` | Cached for periodic re-validation |
| `IsOnline` | `bool` | Fast in-memory presence flag |
| `LogoutTime` | `time.Time` | Set when the session is closed |

## Methods

| Method | Notes |
|--------|-------|
| `Logout()` | Marks offline, records logout time, closes the socket |
| `Write(v)` | Thread-safe `WriteJSON` wrapper |

## Used By

- `store.ChatStore` keeps the canonical `userId -> *Client` map
- `services.ChatImpl` creates or refreshes the client
- `services.MessageImpl`, `NotifyImpl`, `FollowImpl`, `TokenImpl` all read or mutate runtime state on it
