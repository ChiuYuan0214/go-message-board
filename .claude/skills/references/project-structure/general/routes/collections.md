# Routes: /collections

**File:** `backend/general/routes/collections.go`
**Handler:** `CollectionsHandler`

## Endpoints

| Method | Auth | Params / Body | Response |
|--------|------|---------------|----------|
| GET | Yes | `?page=&size=` | `{status, list: CollectionData[]}` |
| POST | Yes | Body: [WriteCollectionData](../types/WriteCollectionData.md) | `{status}` |
| DELETE | Yes | Body: [WriteCollectionData](../types/WriteCollectionData.md) | `{status}` |
