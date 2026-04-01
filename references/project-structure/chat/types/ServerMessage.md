# ServerMessage

**Category:** transport server-only type
**File:** `backend/chat/types/message.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Event` | `string` | Typically `"error"` |
| `Content` | `string` | Human-readable error / status text |

## Used By

- `routes.ChatHandler` when initial token validation fails
- `services.ChatImpl` when a request claims the wrong `userId`
- `services.TokenImpl` when the token checker invalidates a live session
