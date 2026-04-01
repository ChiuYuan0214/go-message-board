# Redis.SetToken

**File:** `backend/security/infra/cache.go`
**Struct:** `Redis`

```go
func (r *Redis) SetToken(userId uint64, token types.Token) error
```

- Redis helper used by security auth flows for token persistence.
