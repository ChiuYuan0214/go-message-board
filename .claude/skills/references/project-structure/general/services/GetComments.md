# GetComments

**File:** `backend/general/services/comments.go`

## Signature

```go
func GetComments(articleId uint64) []types.CommentListData
```

## Returns

All comments for the article as [CommentListData](../types/CommentListData.md). Empty slice on error.

## Behaviour

Fetches all comments with commenter info and aggregated vote counts. No pagination — returns all.
