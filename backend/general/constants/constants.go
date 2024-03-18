package constants

import (
	"os"
)

var (
	PORT = "8080"

	MYSQL_IP       = "your-mysql-address"
	REDIS_IP       = "your-redis-address"
	MYSQL_PASSWORD = "your-mysql-password"
	REDIS_PASSWORD = "your-redis-password"

	JWT_HS256_SECRET_KEY  = "your-jwt-key"
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
