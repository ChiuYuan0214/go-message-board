# CollectionImpl.GetByUserId

**File:** `backend/general/repo/collection.go`
**Struct:** `CollectionImpl`

```go
func (r *CollectionImpl) GetByUserId(userId uint64, start, size int64) (data []types.CollectionData, err error)
```

- Repo method on `CollectionImpl` for the `collection` persistence path in general service.
