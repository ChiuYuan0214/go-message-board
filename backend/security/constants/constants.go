package constants

import (
	"log"
	"os"
	"strconv"
)

var (
	IP_ADDRESS           = "127.0.0.1"
	PORT                 = "7080"
	ADMIN_MAIL_ADDRESS   = "test@gmail.com"
	ADMIN_MAIL_PASSWORD  = "1111"
	JWT_HS256_SECRET_KEY = "jwt_key"

	MYSQL_PORT     = "3306"
	MYSQL_PASSWORD = "1111"
	REDIS_PORT     = "6379"
	REDIS_PASSWORD = "1111"

	VERI_CODE_TTL = 3 // minutes
)

func InitEnv() {
	setVal(&IP_ADDRESS, "IP_ADDRESS")
	setVal(&PORT, "PORT")
	setVal(&ADMIN_MAIL_ADDRESS, "ADMIN_MAIL_ADDRESS")
	setVal(&ADMIN_MAIL_PASSWORD, "ADMIN_MAIL_PASSWORD")
	setVal(&JWT_HS256_SECRET_KEY, "JWT_HS256_SECRET_KEY")
	setVal(&MYSQL_PORT, "MYSQL_PORT")
	setVal(&MYSQL_PASSWORD, "MYSQL_PASSWORD")
	setVal(&REDIS_PORT, "REDIS_PORT")
	setVal(&REDIS_PASSWORD, "REDIS_PASSWORD")
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
