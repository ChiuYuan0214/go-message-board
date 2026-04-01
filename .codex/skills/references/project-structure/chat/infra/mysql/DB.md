# MySQL.DB

**File:** `backend/chat/infra/mysql.go`
**Struct:** `MySQL`

```go
func (m *MySQL) DB() *sql.DB
```

- Returns the raw `*sql.DB` used by repo implementations
- Keeps SQL access behind injected infra instead of package-level globals
- Only repo code should need this handle directly
