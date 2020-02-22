package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./page"))))

	port := "8080"
	log.Printf("start listen on port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
