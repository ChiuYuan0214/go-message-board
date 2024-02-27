package main

import (
	"chat/constants"
	"chat/jobs"
	"chat/routes"
	"chat/services"
	"chat/setup"
	"log"
	"net/http"
)

type CustomLogger struct{}

func main() {
	constants.InitEnv()

	db := setup.InitMySQL()
	defer db.Close()
	services.UsePool(db)

	cache := setup.InitCache()
	defer cache.Client.Close()
	routes.UseCache(cache)
	services.UseCache(cache)

	mongo := setup.InitMongo()
	defer mongo.Close()
	services.UseMongo(mongo)
	jobs.UseMongo(mongo)

	routes.UseConnection()
	jobs.UseScheduler()

	log.Println("listening on port " + constants.PORT)

	err := http.ListenAndServe(constants.PORT, routes.UseCORS())
	if err != nil {
		log.Println(err)
	}
}
