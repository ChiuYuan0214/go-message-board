# FollowerImpl.GetUsersByFollowerId

**File:** `backend/general/repo/follower.go`
**Struct:** `FollowerImpl`

```go
func (r *FollowerImpl) GetUsersByFollowerId(followerId uint64) (users []types.Follower, err error)
```

- Repo method on `FollowerImpl` for the `follower` persistence path in general service.
