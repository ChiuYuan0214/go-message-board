# AddCollection

**File:** `backend/general/services/collection.go`

## Signature

```go
func AddCollection(userId, articleId uint64) bool
```

## Returns

`true` on success. `false` if insert fails (e.g. duplicate).

## Behaviour

Inserts a `(userId, articleId)` row into `collections`.
