# Vote

**Table:** `votes`
**File:** `backend/general/entities/vote.go`

## Fields

| Field | Type | Column | Notes |
|-------|------|--------|-------|
| `VoteId` | `uint64` | `vote_id` | PK, auto-increment |
| `UserId` | `uint64` | `user_id` | FK → users |
| `SourceId` | `uint64` | `source_id` | ID of voted article or comment |
| `Score` | `int8` | `score` | `1` = up, `-1` = down, `0` = retracted |
| `VoteType` | `string` | `vote_type` | `"article"` or `"comment"` |
| `CreationTime` | `time.Time` | `creation_time` | |
| `UpdateTime` | `time.Time` | `update_time` | auto on update |

## Purpose

Unified vote table for both articles and comments. `VoteType` discriminates the target type.

## Related

- Services: [Vote](../services/Vote.md), [UpdateVote](../services/UpdateVote.md)
