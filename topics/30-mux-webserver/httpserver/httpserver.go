package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/30-mux-webserver/httpserver/muxrouter"
)

func Start() {
	// Mux Router
	mux := muxrouter.CreateMuxRouter()

	// Http server
	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: mux,
	}

	fmt.Println("Server running on port 3000")
	log.Fatal(server.ListenAndServe())

}
