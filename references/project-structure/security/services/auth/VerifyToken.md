# AuthImpl.VerifyToken

**File:** `backend/security/services/login.go`
**Struct:** `AuthImpl`

```go
func (s *AuthImpl) VerifyToken(userId uint64, token string) bool
```

- Service method on `AuthImpl` coordinating `auth` behavior in security service.
