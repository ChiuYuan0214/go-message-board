# RegisterImpl.ScheduleCodeInvalidation

**File:** `backend/security/services/register.go`
**Struct:** `RegisterImpl`

```go
func (s *RegisterImpl) ScheduleCodeInvalidation(codeId int64, veriCode *utils.VerificationCode)
```

- Service method on `RegisterImpl` coordinating `register` behavior in security service.
