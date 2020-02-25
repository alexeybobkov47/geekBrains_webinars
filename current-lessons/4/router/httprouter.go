package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func List(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "You see user list\n")
}

func Get(w http.ResponseWriter, r_ *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "you try to see user %s\n", ps.ByName("id"))
}

/*
curl -v -X PUT -H "Content-Type: application/json" -d '{"foo":"fee"}' http://localhost:8080/users
*/

func Create(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "you try to create new user\n")
}

/*
curl -v -X POST -H "Content-Type: application/json" -d '{"foo":"fee"}' http://localhost:8080/users/egor
*/

func Update(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "you try to update %s\n", ps.ByName("login"))
}

func main() {
	router := httprouter.New()

	router.GET("/", List)
	router.GET("/users", List)
	router.PUT("/users", Create)
	router.GET("/users/:id", Get)
	router.POST("/users/:login", Update)

	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
