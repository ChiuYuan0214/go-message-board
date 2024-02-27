package routes

import (
	"database/sql"
	"net/http"
)

var connPool *sql.DB

func UsePool(db *sql.DB) {
	connPool = db
}

func UseDispatcher() {
	http.HandleFunc("/article", handleArticle)
	http.HandleFunc("/articles", handleArticles)
	http.HandleFunc("/comment", authMiddle(handleComment))
	http.HandleFunc("/comments", handleComments)
	http.HandleFunc("/profile", handleProfile)
	http.HandleFunc("/vote", authMiddle(handleVote))
	http.HandleFunc("/view", handleView)
	http.HandleFunc("/follow", authMiddle(handleFollow))
	http.HandleFunc("/follower", handleFollower)
	http.HandleFunc("/follows", handleFollows)
	http.HandleFunc("/collections", authMiddle(handleCollection))
}
