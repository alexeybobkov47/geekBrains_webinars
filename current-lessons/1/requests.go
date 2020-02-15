package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := postReq(); err != nil {
		log.Fatal(err)
	}
}

const (
	getURL  = "https://golang.org"
	postURL = "https://postman-echo.com/post"
)

func postReq() error {
	resp, err := http.Post(
		postURL, "application/json", strings.NewReader(`{"foo1":"bar1","foo2":"bar2"}`))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getReq() error {
	resp, err := http.Get(getURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("Status: %v;\n StatusCode: %v;\n Header: %v\n", resp.Status, resp.StatusCode, resp.Header)

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
