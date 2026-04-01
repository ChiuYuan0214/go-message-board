# FollowerData

**Purpose:** Request — remove a follower (someone following you)
**File:** `backend/general/types/follower.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `Follower` | `follower` | User ID of the follower to remove |

## Used in

- [RemoveFollower](../services/RemoveFollower.md)
- Route: [DELETE /follower](../routes/follower.md)
