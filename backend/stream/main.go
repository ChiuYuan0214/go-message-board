package main

import (
	"stream/constants"
	"stream/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/live", routes.HandleHLS)
	router.GET("/interact", routes.HandleSocket)

	router.Run(constants.PORT)
}
