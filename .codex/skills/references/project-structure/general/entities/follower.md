# Follower

**Table:** `followers`
**File:** `backend/general/entities/follower.go`

## Fields

| Field | Type | Column | Notes |
|-------|------|--------|-------|
| `UserId` | `uint64` | `user_id` | Composite PK — the user being followed |
| `FollowerId` | `uint64` | `follower_id` | Composite PK — the user who is following |

## Purpose

Records a follow relationship. `(UserId, FollowerId)` = "FollowerId follows UserId".

## Related

- Services: [AddFollow](../services/AddFollow.md), [RemoveFollow](../services/RemoveFollow.md), [RemoveFollower](../services/RemoveFollower.md), [GetFollowers](../services/GetFollowers.md), [GetFollows](../services/GetFollows.md)
