package routes

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func customLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("\n\n\n")
		log.Println("---------------------------")
		log.Println("requested URI:", r.RequestURI)
		log.Print("\n")
		log.Println("headers:", r.Header)
		log.Print("\n")
		log.Println("remote address:", r.RemoteAddr)
		log.Println("---------------------------")
		log.Print("\n\n\n")

		next.ServeHTTP(w, r)
	})
}

func UseCORS() http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	return customLogger(c.Handler(http.DefaultServeMux))
}
