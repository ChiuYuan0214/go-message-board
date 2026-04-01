# UpdateVote

**File:** `backend/general/services/vote.go`

## Signature

```go
func UpdateVote(userId, voteId uint64, score int8) bool
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | Must match the vote's owner |
| `voteId` | `uint64` | Existing vote to update |
| `score` | `int8` | New score (`1`, `-1`, or `0` to retract) |

## Returns

`true` on success, `false` on DB error or ownership mismatch.
