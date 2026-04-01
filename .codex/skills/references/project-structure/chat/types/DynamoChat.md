# DynamoChat

**Category:** persistence type
**File:** `backend/chat/types/mongo.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `SenderId` | `uint64` | Message sender |
| `ReceiverId` | `uint64` | Message receiver |
| `Content` | `string` | Message body |
| `Time` | `time.Time` | Stored timestamp in DynamoDB |

## Used By

- `types.DynamoClient` for scan and batch-write payloads
- `repo.HistoryImpl` as the persistence-facing return type
- `services.HistoryImpl` when translating stored records back into `types.Message`
