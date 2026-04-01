# UserInfoList

**Category:** transport response type
**File:** `backend/chat/types/response.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Event` | `string` | `online-follow-list` or `online-follower-list` |
| `List` | `[]uint64` | Online user IDs only |

## Used By

- `services.FollowListImpl.InitFollowList`
- `services.FollowListImpl.InitFollowerList`
