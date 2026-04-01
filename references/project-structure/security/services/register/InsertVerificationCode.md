# RegisterImpl.InsertVerificationCode

**File:** `backend/security/services/register.go`
**Struct:** `RegisterImpl`

```go
func (s *RegisterImpl) InsertVerificationCode(userId int64, code int32, expireTime time.Time) int64
```

- Service method on `RegisterImpl` coordinating `register` behavior in security service.
