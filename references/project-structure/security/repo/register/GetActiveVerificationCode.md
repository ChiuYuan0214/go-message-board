# RegisterImpl.GetActiveVerificationCode

**File:** `backend/security/repo/register.go`
**Struct:** `RegisterImpl`

```go
func (r *RegisterImpl) GetActiveVerificationCode(userId uint64) (code string, err error)
```

- Repo method on `RegisterImpl` for the `register` data path in security service.
