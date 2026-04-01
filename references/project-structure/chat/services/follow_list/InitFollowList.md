# FollowListImpl.InitFollowList

**File:** `backend/chat/services/follow-list.go`
**Struct:** `FollowListImpl`

```go
func (s *FollowListImpl) InitFollowList(wg *sync.WaitGroup, userId uint64)
```

- Loads follow ids via `repo.Follow`
- Stores them on the in-memory client
- Sends `online-follow-list` with the currently online subset
- Used only during connection bootstrap
