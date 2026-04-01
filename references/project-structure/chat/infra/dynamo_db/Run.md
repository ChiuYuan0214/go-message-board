# DynamoDB.Run

**File:** `backend/chat/infra/dynamo.go`
**Struct:** `DynamoDB`

```go
func (d *DynamoDB) Run() (err error)
```

- Creates the AWS session for chat history storage
- Builds the typed DynamoDB client used by history persistence
- This is the depin-managed initialization point for DynamoDB
