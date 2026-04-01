# Client

**Category:** runtime connection state
**File:** `backend/chat/types/chat.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `StateLock` | `sync.RWMutex` | Guards mutable client session state |
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
| `InitSession(conn, token)` | Updates the active socket and token for a live session |
| `ReplaceFollowerList(list)` / `ReplaceFollowList(list)` | Replace in-memory follow graph snapshots |
| `SnapshotFollowerList()` / `SnapshotFollowList()` | Read follow graph snapshots safely |
| `AddFollow(targetUserId)` / `RemoveFollow(targetUserId)` / `RemoveFollower(targetUserId)` | Mutate follow graph state under the client lock |
| `TokenValue()` / `IsActive()` / `LogoutSnapshot()` | Safe readers for shared session state |
| `CloseSession()` | Marks offline, records logout time, closes the socket |
| `Read(v)` | Reads JSON from the current socket |
| `Write(v)` | Thread-safe `WriteJSON` wrapper |

## Used By

- `store.ChatStore` keeps the canonical `userId -> *Client` map
- `services.ChatImpl` creates or refreshes the client session
- `services.MessageImpl`, `NotifyImpl`, `FollowImpl`, `FollowListImpl`, `TokenImpl`, and `jobs.SchedulerImpl` all read or mutate runtime state on it

## Notes

- `Client` is now the main synchronization boundary for per-user chat session state; callers should prefer these methods over touching shared fields directly
