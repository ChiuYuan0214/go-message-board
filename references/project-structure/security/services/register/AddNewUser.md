# RegisterImpl.AddNewUser

**File:** `backend/security/services/register.go`
**Struct:** `RegisterImpl`

```go
func (s *RegisterImpl) AddNewUser(username string, email string, password string, phone string, job string, address string) int64
```

- Service method on `RegisterImpl` coordinating `register` behavior in security service.
