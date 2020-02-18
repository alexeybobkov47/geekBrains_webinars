package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", firstHandle)
	router.HandleFunc("/user", helloUserHandler)
	router.HandleFunc("/examples", examplesHandler)

	port := "8080"
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func firstHandle(wr http.ResponseWriter, _ *http.Request) {
	_, _ = wr.Write([]byte(`Hello, World!`))
}

func helloUserHandler(wr http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(wr, "Hello, %s", req.URL.Query().Get("name"))
}

func examplesHandler(wr http.ResponseWriter, req *http.Request) {
	http.SetCookie(wr, &http.Cookie{
		Name:    "GB",
		Value:   "work",
		Expires: time.Now().Add(time.Minute * 10),
	})

	_, _ = fmt.Fprintf(wr, "Header, %s", req.Header.Get("User-Agent"))

	wr.Header().Set("GeekBrains", "newHeader")
}
