# AuthImpl.Login

**File:** `backend/security/services/login.go`
**Struct:** `AuthImpl`

```go
func (s *AuthImpl) Login(email string, password string) (userId uint64, token *types.Token)
```

- Service method on `AuthImpl` coordinating `auth` behavior in security service.
