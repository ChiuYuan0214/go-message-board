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
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/verifyCode", handleVerifyCode)
	http.HandleFunc("/resendVerificationCode", handleResendCode)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/updatePassword", authMiddle(handleUpdatePassword))
	http.HandleFunc("/updateProfile", authMiddle(handleUpdateProfile))
	http.HandleFunc("/uploadImage", authMiddle(handleUploadImage))

	fs := http.FileServer(http.Dir("./uploads/images/"))
	http.Handle("/uploads/images/", http.StripPrefix("/uploads/images/", fs))
}
