# RemoveFollower

**File:** `backend/general/services/follower.go`

## Signature

```go
func RemoveFollower(userId uint64, follower uint64) bool
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | The user removing a follower |
| `follower` | `uint64` | The follower being removed |

## Returns

`true` on success, `false` on DB error.
