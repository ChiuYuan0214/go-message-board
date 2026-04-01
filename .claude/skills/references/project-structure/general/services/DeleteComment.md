# DeleteComment

**File:** `backend/general/services/comment.go`

## Signature

```go
func DeleteComment(userId uint64, commentId string) (string, int)
```

## Returns

`("", 0)` on success. `(errorMessage, httpStatus)` on failure.

## Behaviour

Validates ownership, then deletes the comment and any related votes from `votes`.
