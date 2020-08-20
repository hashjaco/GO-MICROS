package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string  `json:"last_name"`
	Admin bool `json:"admin"`
}

type Users []User

func allUsers(w http.ResponseWriter, r *http.Request){
	users := Users{
		User{FirstName:"Test User", LastName: "Test Last Name", Admin: false},
	}

	fmt.Println("Endpoint Hit: All Users Endpoint")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}
}

func postUser(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Test Post endpoint works")
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/users", postUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}


func main() {
	handleRequests()
}
