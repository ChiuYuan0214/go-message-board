# UpdateComment

**File:** `backend/general/services/comment.go`

## Signature

```go
func UpdateComment(userId uint64, data *types.UpdateCommentData) (string, int)
```

## Returns

`("", 0)` on success. `(errorMessage, httpStatus)` on failure.

## Behaviour

- Returns `400` if comment not found
- Returns `400` if `userId` doesn't own the comment
- Partial update: only non-empty fields are written
