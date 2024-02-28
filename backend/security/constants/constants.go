package constants

import (
	"log"
	"os"
	"strconv"
)

var (
	PORT = "7080"

	MYSQL_IP       = "127.0.0.1:3306"
	REDIS_IP       = "127.0.0.1:6379"
	MYSQL_PASSWORD = "1111"
	REDIS_PASSWORD = "test1234"

	ADMIN_MAIL_ADDRESS   = "test@gmail.com"
	ADMIN_MAIL_PASSWORD  = "1111"
	JWT_HS256_SECRET_KEY = "jwt_key"

	VERI_CODE_TTL = 3 // minutes
)

func InitEnv() {
	setVal(&PORT, "PORT")
	setVal(&MYSQL_IP, "MYSQL_IP")
	setVal(&REDIS_IP, "REDIS_IP")
	setVal(&MYSQL_PASSWORD, "MYSQL_PASSWORD")
	setVal(&REDIS_PASSWORD, "REDIS_PASSWORD")
	setVal(&ADMIN_MAIL_ADDRESS, "ADMIN_MAIL_ADDRESS")
	setVal(&ADMIN_MAIL_PASSWORD, "ADMIN_MAIL_PASSWORD")
	setVal(&JWT_HS256_SECRET_KEY, "JWT_HS256_SECRET_KEY")
	setIntVal(&VERI_CODE_TTL, "VERI_CODE_TTL")
}

func setVal(key *string, env string) {
	val := os.Getenv(env)
	if val != "" {
		*key = val
	}
}

func setIntVal(key *int, env string) {
	val := os.Getenv(env)
	if val != "" {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Println(err)
			return
		}
		*key = num
	}
}
