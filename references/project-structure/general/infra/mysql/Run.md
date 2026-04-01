# MySQL.Run

**File:** `backend/general/infra/mysql.go`
**Struct:** `MySQL`

```go
func (i *MySQL) Run() (err error)
```

- Initializes the GORM MySQL connection for the general service.
- Open this when changing DB bootstrap, DSN, or retry behavior.
