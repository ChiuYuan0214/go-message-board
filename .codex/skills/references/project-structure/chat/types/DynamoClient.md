# DynamoClient

**Category:** persistence helper
**File:** `backend/chat/types/mongo.go`

## Methods

| Method | Notes |
|--------|-------|
| `GetAllWithFilters(senderId, receiverId, startTime, endTime)` | Reads history rows for one conversation window |
| `BatchInsert(chatList)` | Writes cached chat rows in bulk |

## Used By

- `infra.DynamoDB.Client()` returns this wrapper
- `repo.HistoryImpl` delegates all DynamoDB access to it

## Notes

- This is part of the chat history persistence boundary, not a general-purpose Dynamo helper
