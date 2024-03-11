package jobs

import "database/sql"

var connPool *sql.DB

func UsePool(db *sql.DB) {
	connPool = db
}

func UseJobs() {
	go initUsers()
}
