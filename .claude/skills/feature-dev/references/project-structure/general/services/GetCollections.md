# GetCollections

**File:** `backend/general/services/collection.go`

## Signature

```go
func GetCollections(userId uint64, page, size int64) []types.CollectionData
```

## Returns

Paginated array of [CollectionData](../types/CollectionData.md). Empty slice on error.

## Behaviour

Returns articles the user has collected, with full article detail including vote counts. Sorted by collection creation time DESC.
