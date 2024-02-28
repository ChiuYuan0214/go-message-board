package constants

import (
	"os"
)

var (
	PORT = "8080"

	MYSQL_IP       = "127.0.0.1:3306"
	REDIS_IP       = "127.0.0.1:6379"
	MYSQL_PASSWORD = "1111"
	REDIS_PASSWORD = "test1234"

	JWT_HS256_SECRET_KEY  = "jwt-key"
	VIEW_COUNT_CACHE_NAME = "viewCount"
	VIEW_LIST_NAME        = "viewList"
	COMMENTED_ARTICLE_SET = "commentedArticleSet"
	HOT_LIST_NAMAE        = "hotList"
)

func InitEnv() {
	setVal(&PORT, "PORT")
	setVal(&MYSQL_IP, "MYSQL_IP")
	setVal(&REDIS_IP, "REDIS_IP")
	setVal(&MYSQL_PASSWORD, "MYSQL_PASSWORD")
	setVal(&REDIS_PASSWORD, "REDIS_PASSWORD")
	setVal(&JWT_HS256_SECRET_KEY, "JWT_HS256_SECRET_KEY")
}

func setVal(key *string, env string) {
	val := os.Getenv(env)
	if val != "" {
		*key = val
	}
}
