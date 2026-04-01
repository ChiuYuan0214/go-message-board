# TokenImpl.GetToken

**File:** `backend/chat/repo/token.go`
**Struct:** `TokenImpl`

```go
func (r *TokenImpl) GetToken(userId uint64) (string, error)
```

- Reads the current token for a user through `infra.Cache`
- Backing lookup used by `services.TokenImpl.ValidateToken`
