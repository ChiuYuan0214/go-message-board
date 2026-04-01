# TokenImpl.UseTokenChecker

**File:** `backend/chat/services/token.go`
**Struct:** `TokenImpl`

```go
func (s *TokenImpl) UseTokenChecker(ctx context.Context, cancel context.CancelFunc, userId uint64)
```

- Re-validates the connected client's token every 10 minutes
- Sends a token error and cancels the session if validation fails
- Runs in its own goroutine per active connection
