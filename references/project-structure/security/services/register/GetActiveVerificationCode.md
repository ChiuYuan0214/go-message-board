# RegisterImpl.GetActiveVerificationCode

**File:** `backend/security/services/register.go`
**Struct:** `RegisterImpl`

```go
func (s *RegisterImpl) GetActiveVerificationCode(userId uint64) (string, error)
```

- Service method on `RegisterImpl` coordinating `register` behavior in security service.
