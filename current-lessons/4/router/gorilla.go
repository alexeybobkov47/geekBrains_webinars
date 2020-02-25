package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ListMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You see user list\n")
}

func GetMux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to see user %s\n", vars["id"])
}

/*
curl -v -X PUT -H "Content-Type: application/json" -d '{"foo":"fee"}' http://localhost:8080/users
*/

func CreateMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you try to create new user\n")
}

/*
curl -v -X POST -H "Content-Type: application/json"  -H "X-Auth: test" -d '{"foo":"fee"}' http://localhost:8080/users/login
*/

func UpdateMux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to update %s\n", vars["login"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ListMux)

	r.HandleFunc("/users", ListMux).Host("localhost")

	r.HandleFunc("/users", UpdateMux).Methods("PUT")

	r.HandleFunc("/users/{id:[0-9]+}", GetMux)

	r.HandleFunc("/users/{login}", CreateMux).Methods("POST").Headers("X-Auth", "test")

	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
