# Notification

**Category:** transport event type
**File:** `backend/chat/types/response.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Event` | `string` | `follow-login`, `follow-logout`, `follower-login`, or `follower-logout` |
| `UserId` | `uint64` | User whose online status changed |

## Used By

- `services.NotifyImpl` for presence notifications
- `services.FollowImpl.AddFollow` for immediate follow-login feedback
