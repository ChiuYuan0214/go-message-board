# FollowerImpl.GetByUserIdAndFollowerId

**File:** `backend/general/repo/follower.go`
**Struct:** `FollowerImpl`

```go
func (r *FollowerImpl) GetByUserIdAndFollowerId(userId, followerId uint64) (follower entities.Follower, err error)
```

- Repo method on `FollowerImpl` for the `follower` persistence path in general service.
