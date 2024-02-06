package setup

import (
	"fmt"
	"general/constants"
	"log"
	"net/http"
)

func InitServer() {
	fmt.Printf("Server listening on :%s...\n", constants.PORT)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", constants.PORT), nil))
}
