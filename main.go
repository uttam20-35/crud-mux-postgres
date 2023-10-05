package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	// Initialize Logrus with desired configurations
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel) // Set log level to display INFO and above
	log.Info("Starting the Application")
	InitialMigration()
	initializeRouter()
}
