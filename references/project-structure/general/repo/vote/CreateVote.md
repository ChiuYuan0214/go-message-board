# VoteImpl.CreateVote

**File:** `backend/general/repo/vote.go`
**Struct:** `VoteImpl`

```go
func (r *VoteImpl) CreateVote(userId, sourceId uint64, score int8, voteType string) (id uint64, err error)
```

- Repo method on `VoteImpl` for the `vote` persistence path in general service.
