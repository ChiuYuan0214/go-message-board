# ProfileImpl.UpdatePassword

**File:** `backend/security/repo/profile.go`
**Struct:** `ProfileImpl`

```go
func (r *ProfileImpl) UpdatePassword(userId uint64, password string) (int64, error)
```

- Repo method on `ProfileImpl` for the `profile` data path in security service.
