package constants

import (
	"os"
)

var (
	IP_ADDRESS            = "127.0.0.1"
	PORT                  = "8080"
	MYSQL_PORT            = "3306"
	MYSQL_PASSWORD        = "1111"
	REDIS_PORT            = "6379"
	REDIS_PASSWORD        = "test1234"
	JWT_HS256_SECRET_KEY  = "jwt-key"
	VIEW_COUNT_CACHE_NAME = "viewCount"
	VIEW_LIST_NAME        = "viewList"
	COMMENTED_ARTICLE_SET = "commentedArticleSet"
	HOT_LIST_NAMAE        = "hotList"
)

func InitEnv() {
	setVal(&IP_ADDRESS, "IP_ADDRESS")
	setVal(&PORT, "PORT")
	setVal(&MYSQL_PORT, "MYSQL_PORT")
	setVal(&MYSQL_PASSWORD, "MYSQL_PASSWORD")
	setVal(&REDIS_PORT, "REDIS_PORT")
	setVal(&REDIS_PASSWORD, "REDIS_PASSWORD")
	setVal(&JWT_HS256_SECRET_KEY, "JWT_HS256_SECRET_KEY")
}

func setVal(key *string, env string) {
	val := os.Getenv(env)
	if val != "" {
		*key = val
	}
}
