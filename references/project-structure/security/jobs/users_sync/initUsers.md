# UsersSyncJob.initUsers

**File:** `backend/security/jobs/init-users.go`
**Struct:** `UsersSyncJob`

```go
func (j *UsersSyncJob) initUsers()
```

- Loads users from repo and refreshes the shared in-memory directory on schedule.
