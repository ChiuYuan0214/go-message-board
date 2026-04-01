# UseTokenChecker

**File:** `backend/chat/services/`

## Signature

```go
func UseTokenChecker(ctx context.Context, cancel context.CancelFunc, userId uint64)
```

## Behaviour

Goroutine that runs on a 10-minute ticker. Re-validates the client's token against Redis.
If invalid, sends an error message to the client and calls `cancel()` to terminate the session.
