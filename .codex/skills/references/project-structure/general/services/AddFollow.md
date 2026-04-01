# AddFollow

**File:** `backend/general/services/follow.go`

## Signature

```go
func AddFollow(userId uint64, followee uint64) bool
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | The user who is following |
| `followee` | `uint64` | The user being followed |

## Returns

`true` on success. Idempotent — re-following returns `true` without error.
