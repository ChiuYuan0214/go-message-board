# FollowerImpl.GetFollowersByUserId

**File:** `backend/general/repo/follower.go`
**Struct:** `FollowerImpl`

```go
func (r *FollowerImpl) GetFollowersByUserId(userId uint64) (followers []types.Follower, err error)
```

- Repo method on `FollowerImpl` for the `follower` persistence path in general service.
