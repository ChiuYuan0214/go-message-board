# VoteImpl.GetByUserAndSource

**File:** `backend/general/repo/vote.go`
**Struct:** `VoteImpl`

```go
func (r *VoteImpl) GetByUserAndSource(userId, sourceId uint64) (vote entities.Vote, err error)
```

- Repo method on `VoteImpl` for the `vote` persistence path in general service.
