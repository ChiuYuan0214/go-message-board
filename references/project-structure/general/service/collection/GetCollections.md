# CollectionImpl.GetCollections

**File:** `backend/general/service/collection.go`
**Struct:** `CollectionImpl`

```go
func (s *CollectionImpl) GetCollections(userId uint64, page, size int64) (data []types.CollectionData)
```

- Service method on `CollectionImpl` coordinating `collection` business logic in general service.
