# VoteImpl.Vote

**File:** `backend/general/service/vote.go`
**Struct:** `VoteImpl`

```go
func (s *VoteImpl) Vote(userId, sourceId uint64, score int8, voteType *string) (string, uint64)
```

- Service method on `VoteImpl` coordinating `vote` business logic in general service.
