# TokenImpl.ValidateToken

**File:** `backend/chat/services/token.go`
**Struct:** `TokenImpl`

```go
func (s *TokenImpl) ValidateToken(token string, userId uint64) bool
```

- Looks up the expected token through `repo.Token`
- Returns `true` only when the provided token matches Redis state
- Used both at initial WebSocket connect time and by the periodic token checker
