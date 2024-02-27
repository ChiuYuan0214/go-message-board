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
	defer db.Close()
	routes.UsePool(db)
	services.UsePool(db)
	jobs.UsePool(db)

	cache := setup.InitCache()
	defer cache.Client.Close()
	services.UseCache(cache)
	jobs.UseCache(cache)

	routes.UseDispatcher()
	jobs.UseScheduler()

	setup.InitServer()
}
