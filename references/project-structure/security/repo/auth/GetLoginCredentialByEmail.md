# AuthImpl.GetLoginCredentialByEmail

**File:** `backend/security/repo/auth.go`
**Struct:** `AuthImpl`

```go
func (r *AuthImpl) GetLoginCredentialByEmail(email string) (userId uint64, hashedPassword string, err error)
```

- Repo method on `AuthImpl` for the `auth` data path in security service.
