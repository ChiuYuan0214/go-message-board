# FollowData

**Purpose:** Request — follow or unfollow a user
**File:** `backend/general/types/follower.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `Followee` | `followee` | User ID of person to follow/unfollow |

## Used in

- [AddFollow](../services/AddFollow.md), [RemoveFollow](../services/RemoveFollow.md)
- Route: [POST/DELETE /follow](../routes/follow.md)
