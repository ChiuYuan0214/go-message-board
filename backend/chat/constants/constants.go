package constants

import "os"

var (
	PORT = ":9080"

	MYSQL_IP                   = "127.0.0.1:3306"
	REDIS_IP                   = "127.0.0.1:6379"
	MONGO_IP                   = "127.0.0.1:27017"
	MYSQL_PASSWORD             = "1111"
	REDIS_PASSWORD             = "test"
	MONGO_INITDB_DATABASE      = "go-local-database"
	MONGO_INITDB_ROOT_USERNAME = "tester"
	MONGO_INITDB_ROOT_PASSWORD = "1111"

	JWT_HS256_SECRET_KEY = "jwt-secret"
)

func InitEnv() {
	setVal(&PORT, "PORT")

	setVal(&MYSQL_IP, "MYSQL_IP")
	setVal(&REDIS_IP, "REDIS_IP")
	setVal(&MONGO_IP, "MONGO_IP")
	setVal(&MYSQL_PASSWORD, "MYSQL_PASSWORD")
	setVal(&REDIS_PASSWORD, "REDIS_PASSWORD")
	setVal(&MONGO_INITDB_DATABASE, "MONGO_INITDB_DATABASE")
	setVal(&MONGO_INITDB_ROOT_USERNAME, "MONGO_INITDB_ROOT_USERNAME")
	setVal(&MONGO_INITDB_ROOT_PASSWORD, "MONGO_INITDB_ROOT_PASSWORD")

	setVal(&JWT_HS256_SECRET_KEY, "JWT_HS256_SECRET_KEY")
}

func setVal(key *string, env string) {
	val := os.Getenv(env)
	if val != "" {
		*key = val
	}
}
