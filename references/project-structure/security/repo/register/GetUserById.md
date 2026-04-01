# RegisterImpl.GetUserById

**File:** `backend/security/repo/register.go`
**Struct:** `RegisterImpl`

```go
func (r *RegisterImpl) GetUserById(userId uint64) (user store.User, err error)
```

- Repo method on `RegisterImpl` for the `register` data path in security service.
