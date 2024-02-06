package main

import (
	"general/constants"
	"general/jobs"
	"general/routes"
	"general/services"
	"general/setup"
)

const PORT = 8080

func main() {
	constants.InitEnv()
	db := setup.InitMySQL()
	routes.UsePool(db)
	services.UsePool(db)
	jobs.UsePool(db)

	cache := setup.InitCache()
	services.UseCache(cache)
	jobs.UseCache(cache)

	routes.UseDispatcher()
	jobs.UseScheduler()

	setup.InitServer()
	db.Close()
}
