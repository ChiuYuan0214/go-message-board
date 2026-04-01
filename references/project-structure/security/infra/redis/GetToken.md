# Redis.GetToken

**File:** `backend/security/infra/cache.go`
**Struct:** `Redis`

```go
func (r *Redis) GetToken(userId uint64) (string, error)
```

- Redis helper used by security auth flows for token persistence.
