# GetFollowers

**File:** `backend/general/services/follower.go`

## Signature

```go
func GetFollowers(userId uint64) []types.Follower
```

## Returns

All users following `userId` as [Follower](../types/Follower.md). Empty slice on error.
