# RequestEvent

**Category:** transport request type
**File:** `backend/chat/types/request.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `UserId` | `uint64` | Sender |
| `TargetUserId` | `uint64` | Target user for message / history / follow events |
| `Type` | `string` | Event type such as `message`, `history`, `add-follow`, `remove-follow`, `remove-follower`, `refresh-token`, `ping` |
| `Content` | `string` | Message body or history cursor timestamp, depending on event type |

## Used By

- Read from the socket by `services.ChatImpl.ListenChatEvent`
- Routed by `services.EventImpl.HandleEvent`
- Consumed by `MessageImpl`, `HistoryImpl`, and `FollowImpl`
