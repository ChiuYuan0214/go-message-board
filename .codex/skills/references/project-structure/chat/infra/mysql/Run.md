# MySQL.Run

**File:** `backend/chat/infra/mysql.go`
**Struct:** `MySQL`

```go
func (m *MySQL) Run() (err error)
```

- Opens the chat service MySQL connection
- Retries until MySQL is reachable
- Applies pool settings before the app starts serving traffic
- This retry loop makes chat startup wait for MySQL instead of failing immediately
