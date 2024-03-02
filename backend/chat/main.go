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

	dynamo := setup.InitDynamo()
	services.UseDynamo(dynamo)
	jobs.UseDynamo(dynamo)

	routes.UseConnection()
	jobs.UseScheduler()

	log.Println("listening on port " + constants.PORT)

	err := http.ListenAndServe(constants.PORT, routes.UseCORS())
	if err != nil {
		log.Println(err)
	}
}
