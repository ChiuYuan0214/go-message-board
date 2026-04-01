# DynamoDB.Client

**File:** `backend/chat/infra/dynamo.go`
**Struct:** `DynamoDB`

```go
func (d *DynamoDB) Client() *types.DynamoClient
```

- Returns the typed DynamoDB client wrapper
- Used by history repo code for read and batch-write operations
- Open this together with `types/DynamoClient.md` when changing history persistence behavior
