package handlers

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/32-api-rest/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	// Get Users
	users := models.ListUsers()

	// Response
	SendData(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Obtain id from URL parameter
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get data
	user := models.GetUser(id)

	// Response
	if user.Id == 0 {
		SendNotFound(w)
	} else {
		SendData(w, user)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Create user model
	user := models.User{}

	// Create decoder with request body data
	decoder := json.NewDecoder(r.Body)

	// Decode and save data
	if err := decoder.Decode(&user); err != nil {
		SendUnprocessable(w)
	} else {
		user.Save()
		SendData(w, user)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Obtain id from URL parameter
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get data
	user := models.GetUser(id)

	// Create decoder with request body data
	decoder := json.NewDecoder(r.Body)

	// Decode and save data
	if err := decoder.Decode(&user); err != nil {
		SendUnprocessable(w)
	} else {
		SendData(w, user)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Obtain id from URL parameter
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Delete user
	user := models.GetUser(id)
	user.Delete()

	SendData(w, user)
}
