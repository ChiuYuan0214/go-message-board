# ProfileImpl.GetPasswordByUserId

**File:** `backend/security/repo/profile.go`
**Struct:** `ProfileImpl`

```go
func (r *ProfileImpl) GetPasswordByUserId(userId uint64) (hashedPassword string, err error)
```

- Repo method on `ProfileImpl` for the `profile` data path in security service.
