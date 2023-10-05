package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm" //this package or module for the GORM implementation.
)

var DB *gorm.DB
var err error

const DNS = "postgres://postgres:root@localhost:5432/go_crud_rest_api" //Database Configuration

type User struct {
	gorm.Model //no need to create manual id cause "gorm.Model" contains all this stuff like id,created/updated/deleted timestamp

	// to convert this struct into model, use "GORM"
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// this function includes all the database configuration
func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB")
	}
	DB.AutoMigrate(&User{})
}

// Returning List of Users available
func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("List Of Users")
	w.Header().Set("Content-Type", "application/json")
	var users []User //struct of users
	DB.Find(&users)  //giving only ref
	json.NewEncoder(w).Encode(users)
}

// Returning particular user matched with id
func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Get an User with regarding Id")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User //struct of users
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

// Creates a new user with its data
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating an user")
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// Update particular user by its ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating an user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// Delete the user by its ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting an user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User is deleted")
}
