# Routes: /follower

**File:** `backend/general/routes/follower.go`
**Handler:** `FollowerHandler`

## Endpoints

| Method | Auth | Params / Body | Response |
|--------|------|---------------|----------|
| GET | No | `?userId=` | `{status, list: Follower[]}` |
| DELETE | Yes | `{follower: uint64}` | `{status}` |

## Notes

GET returns followers of any user. DELETE removes a follower from the authenticated user's follower list.
