# ChatStore

**File:** `backend/chat/store/chat.go`

## Purpose

- Owns the in-memory client map and shared broadcast channel
- Creates clients lazily through `GetClient`
- Exposes helper methods such as `DeleteClient` and `GetSendMap`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `mu` | `sync.RWMutex` | Guards the client map |
| `clients` | `map[uint64]*types.Client` | Canonical in-memory presence map |
| `Broadcast` | `chan *types.RequestEvent` | Shared event channel consumed by `EventImpl` |

## Key Methods

| Method | Notes |
|--------|-------|
| `CreateClient(userId)` | Creates a minimal client shell; used under store lock |
| `GetClient(userId)` | Returns a client and lazily creates one if missing |
| `FindClient(userId)` | Reads an existing client without creating a placeholder |
| `SnapshotClients()` | Returns a shallow snapshot for safe iteration outside the store lock |
| `DeleteClient(userId)` | Removes a client from the map |
| `GetSendMap(userId)` | Returns the client's per-target message cache |

## Notes

- This store is still package-level singleton state exposed through `GetChatStore()`
- The client map is now guarded by an internal `RWMutex`; callers should prefer store methods over direct map access
- `GetClient()` is the create-if-missing path; `FindClient()` is the read-only lookup path for existing sessions
- Any change to client lifecycle or event fan-out should be checked against this file
