# ChatStore

**File:** `backend/chat/store/chat.go`

## Purpose

- Owns the in-memory client map and shared broadcast channel
- Creates clients lazily through `GetClient`
- Exposes helper methods such as `DeleteClient` and `GetSendMap`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Clients` | `*map[uint64]*types.Client` | Canonical in-memory presence map |
| `Broadcast` | `chan *types.RequestEvent` | Shared event channel consumed by `EventImpl` |

## Key Methods

| Method | Notes |
|--------|-------|
| `CreateClient(userId)` | Creates a minimal client shell |
| `GetClient(userId)` | Returns a client and lazily creates one if missing |
| `DeleteClient(userId)` | Removes a client from the map |
| `GetSendMap(userId)` | Returns the client's per-target message cache |

## Notes

- This store is still package-level singleton state exposed through `GetChatStore()`
- Any change to client lifecycle or event fan-out should be checked against this file
