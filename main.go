package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	r := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	r.HandleFunc("/createProfile", createProfile).Methods("POST")
	r.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")
	r.HandleFunc("/getUserProfile/{_id}", getUserProfile).Methods("GET")
	r.HandleFunc("/updateProfileName/{_id}", updateProfileName).Methods("PUT")
	r.HandleFunc("/deleteProfile/{_id}", deleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r)) // Run Server
}
