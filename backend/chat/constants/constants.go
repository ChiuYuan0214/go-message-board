package constants

import "os"

var (
	IP_ADDRESS                 = "127.0.0.1"
	MONGO_INITDB_DATABASE      = "go-local-database"
	MONGO_INITDB_ROOT_USERNAME = "tester"
	MONGO_INITDB_ROOT_PASSWORD = "1111"
	MONGO_PORT                 = "27017"
	PORT                       = ":9080"
	JWT_HS256_SECRET_KEY       = "jwt-secret"
	MYSQL_PORT                 = "3306"
	MYSQL_PASSWORD             = "1111"
	REDIS_PORT                 = "6379"
	REDIS_PASSWORD             = "test"
)

func InitEnv() {
	setVal(&IP_ADDRESS, "IP_ADDRESS")
	setVal(&MONGO_INITDB_DATABASE, "MONGO_INITDB_DATABASE")
	setVal(&MONGO_INITDB_ROOT_USERNAME, "MONGO_INITDB_ROOT_USERNAME")
	setVal(&MONGO_INITDB_ROOT_PASSWORD, "MONGO_INITDB_ROOT_PASSWORD")
	setVal(&MONGO_PORT, "MONGO_PORT")
	setVal(&PORT, "PORT")
	setVal(&JWT_HS256_SECRET_KEY, "JWT_HS256_SECRET_KEY")
	setVal(&MYSQL_PORT, "MYSQL_PORT")
	setVal(&MYSQL_PASSWORD, "MYSQL_PASSWORD")
	setVal(&REDIS_PORT, "REDIS_PORT")
	setVal(&REDIS_PASSWORD, "REDIS_PASSWORD")

}

func setVal(key *string, env string) {
	val := os.Getenv(env)
	if val != "" {
		*key = val
	}
}
