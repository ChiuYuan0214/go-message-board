# ProfileImpl.InsertProfileImageInfo

**File:** `backend/security/services/profile.go`
**Struct:** `ProfileImpl`

```go
func (s *ProfileImpl) InsertProfileImageInfo(userId *uint64, fileName *string, desc *string) (string, int)
```

- Service method on `ProfileImpl` coordinating `profile` behavior in security service.
