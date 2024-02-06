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
	http.HandleFunc("/comment", handleComment)
	http.HandleFunc("/comments", handleComments)
	http.HandleFunc("/profile", handleProfile)
	http.HandleFunc("/vote", authMiddle(handleVote))
	http.HandleFunc("/view", handleView)
}
