# RemoveFollow

**File:** `backend/general/services/follow.go`

## Signature

```go
func RemoveFollow(userId uint64, followee uint64) bool
```

## Returns

`true` on success. Idempotent — unfollowing a non-followed user returns `true`.
