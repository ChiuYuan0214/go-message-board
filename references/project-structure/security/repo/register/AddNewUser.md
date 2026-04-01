# RegisterImpl.AddNewUser

**File:** `backend/security/repo/register.go`
**Struct:** `RegisterImpl`

```go
func (r *RegisterImpl) AddNewUser(username string, email string, password string, phone string, job string, address string) (int64, error)
```

- Repo method on `RegisterImpl` for the `register` data path in security service.
