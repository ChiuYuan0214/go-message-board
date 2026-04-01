# RegisterImpl.InsertVerificationCode

**File:** `backend/security/repo/register.go`
**Struct:** `RegisterImpl`

```go
func (r *RegisterImpl) InsertVerificationCode(userId int64, code int32, expireTime time.Time) (int64, error)
```

- Repo method on `RegisterImpl` for the `register` data path in security service.
