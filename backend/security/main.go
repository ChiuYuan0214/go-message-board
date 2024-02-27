package main

import (
	"security/constants"
	"security/routes"
	"security/services"
	"security/setup"
)

func main() {
	constants.InitEnv()

	db := setup.InitMySQL()
	defer db.Close()
	routes.UsePool(db)
	services.UsePool(db)

	cache := setup.InitCache()
	defer cache.Client.Close()
	services.UseCache(cache)

	routes.UseDispatcher()
	setup.InitServer()
}
