package main

import (
	"log"
	"net/http"
	"text/template"
)

type ChallengeList struct {
	Name        string
	Description string
	List        []Challenge
}

type Challenge struct {
	ID       string
	Text     string
	Complete bool
}

var tmpl = template.Must(template.New("MyTemplate").ParseFiles("taskList.html"))

var simpleList = ChallengeList{
	Name:        "Название листа",
	Description: "Описание листа с задачами",
	List: []Challenge{
		{ID: "first", Text: "Первая задача", Complete: false},
		{ID: "second", Text: "Вторая задача", Complete: false},
		{ID: "thrid", Text: "Третья задача", Complete: true},
	},
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", viewList)

	port := "8080"
	log.Printf("server start at port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func viewList(w http.ResponseWriter, req *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "list", simpleList); err != nil {
		log.Println(err)
	}
}
