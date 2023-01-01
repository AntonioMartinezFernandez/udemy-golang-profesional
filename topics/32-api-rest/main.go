package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/32-api-rest/db"
	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/32-api-rest/handlers"
)

func main() {
	// Connect DB
	db.Connect()

	// Http Server and Router
	gorillaMux := mux.NewRouter()

	gorillaMux.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	gorillaMux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	gorillaMux.HandleFunc("/api/user", handlers.CreateUser).Methods("POST")
	gorillaMux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	gorillaMux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	srv := &http.Server{
		Handler: gorillaMux,
		Addr:    "0.0.0.0:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
