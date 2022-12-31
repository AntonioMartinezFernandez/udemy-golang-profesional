package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/29-webserver/httpserver/router"
)

func Start() {
	// Router
	router.CreateRouter()

	// Http server
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}
