package setup

import (
	"fmt"
	"log"
	"net/http"
	"security/constants"
)

func InitServer() {
	fmt.Printf("Server listening on :%s...\n", constants.PORT)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", constants.PORT), nil))
}
