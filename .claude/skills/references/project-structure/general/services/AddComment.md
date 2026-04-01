# AddComment

**File:** `backend/general/services/comment.go`

## Signature

```go
func AddComment(userId uint64, data *types.AddCommentData) uint64
```

## Returns

New comment ID. `0` on error.

## Behaviour

Validates the target article exists, then inserts a `comments` row. Returns the new `comment_id`.
