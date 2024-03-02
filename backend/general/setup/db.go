package setup

import (
	"database/sql"
	"fmt"
	"general/constants"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL() *sql.DB {
	var pool *sql.DB
	var err error
	for {
		connectionString := "root:" + constants.MYSQL_PASSWORD + "@tcp(" + constants.MYSQL_IP + ")/go_project?parseTime=true"

		// create a connection pool
		pool, err = sql.Open("mysql", connectionString)
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		err = pool.Ping()
		if err != nil {
			log.Println(err)
			pool.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("connected to mysql")
		pool.SetMaxOpenConns(15)
		pool.SetMaxIdleConns(5)
		pool.SetConnMaxIdleTime(time.Minute * 30)

		break
	}

	return pool
}
