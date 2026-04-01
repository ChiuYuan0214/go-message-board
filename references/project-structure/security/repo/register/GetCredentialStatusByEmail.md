# RegisterImpl.GetCredentialStatusByEmail

**File:** `backend/security/repo/register.go`
**Struct:** `RegisterImpl`

```go
func (r *RegisterImpl) GetCredentialStatusByEmail(email string) (userId int64, hashedPassword string, isActive bool, err error)
```

- Repo method on `RegisterImpl` for the `register` data path in security service.
