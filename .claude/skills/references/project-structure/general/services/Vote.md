# Vote

**File:** `backend/general/services/vote.go`

## Signature

```go
func Vote(userId, sourceId uint64, score int8, voteType *string) (string, uint64)
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | Voter |
| `sourceId` | `uint64` | Target article or comment ID |
| `score` | `int8` | `1` = up, `-1` = down |
| `voteType` | `*string` | `"article"` or `"comment"` |

## Returns

| Value | Notes |
|-------|-------|
| `string` | Error message; `""` on success |
| `uint64` | New vote ID; `0` on error |

## Behaviour

Validates the source exists (for `voteType == "article"` checks `articles` table).
Inserts a new vote record. For toggling an existing vote, call [UpdateVote](UpdateVote.md) instead.
