# GetHistory

**File:** `backend/chat/services/`

## Signature

```go
func GetHistory(event *types.RequestEvent)
```

## Behaviour

1. Fetches message history from DynamoDB for the conversation (default 5-hour window)
2. Merges with in-memory `SendMap` cache to include recent unsync'd messages
3. Sends a [History](../types/History.md) response back to the requesting client
