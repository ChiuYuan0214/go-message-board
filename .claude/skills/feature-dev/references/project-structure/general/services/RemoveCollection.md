# RemoveCollection

**File:** `backend/general/services/collection.go`

## Signature

```go
func RemoveCollection(userId, articleId uint64) bool
```

## Returns

`true` on success, `false` on DB error.

## Behaviour

Deletes the `(userId, articleId)` row from `collections`.
