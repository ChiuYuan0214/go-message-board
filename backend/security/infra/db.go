package infra

import (
	"database/sql"
	"fmt"
	"log"
	"security/constants"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var _ RDB = (*MySQL)(nil)

type MySQL struct {
	db *sql.DB
}

func (m *MySQL) Run() (err error) {
	for {
		connectionString := "root:" + constants.MYSQL_PASSWORD + "@tcp(" + constants.MYSQL_IP + ")/go_project?parseTime=true"

		// create a connection pool
		m.db, err = sql.Open("mysql", connectionString)
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		err = m.db.Ping()
		if err != nil {
			log.Println(err)
			m.db.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("connected to mysql")
		m.db.SetMaxOpenConns(15)
		m.db.SetMaxIdleConns(5)
		m.db.SetConnMaxIdleTime(time.Minute * 30)

		break
	}
	return
}

func (m *MySQL) Stop() {
	if m.db != nil {
		m.db.Close()
	}
}

func (m *MySQL) DB() *sql.DB {
	return m.db
}
