# FollowImpl.AddFollow

**File:** `backend/chat/services/follow.go`
**Struct:** `FollowImpl`

```go
func (s *FollowImpl) AddFollow(event *types.RequestEvent)
```

- Adds `TargetUserId` to the caller's in-memory `FollowList`
- Sends `follow-login` immediately if the target is online
- Does not persist follow data; DB ownership stays outside chat
