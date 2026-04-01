# GetFollows

**File:** `backend/general/services/follows.go`

## Signature

```go
func GetFollows(userId uint64) []types.Follower
```

## Returns

All users that `userId` is following, as [Follower](../types/Follower.md). Empty slice on error.
