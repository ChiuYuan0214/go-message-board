# Routes: /vote

**File:** `backend/general/routes/vote.go`
**Handler:** `VoteHandler`

## Endpoints

| Method | Auth | Body | Response |
|--------|------|------|----------|
| POST | Yes | `{sourceId, score, voteType}` | `{status, id}` |
| PUT | Yes | `{voteId, score}` | `{status}` |

## Notes

- `score`: `1` = upvote, `-1` = downvote
- `voteType`: `"article"` or `"comment"`
- POST → [Vote](../services/Vote.md); PUT → [UpdateVote](../services/UpdateVote.md)
