package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var DSN = "root:1234@tcp(localhost:3306)/task_list_app?charset=utf8"

// TaskList - список задач
type TaskList struct {
	ID          int
	Name        string
	Description string
	List        []Task
}

// Task - задача и ее статус
type Task struct {
	ID       string
	ListID   int
	Text     string
	Complete bool
}

var tmpl = template.Must(template.New("MyTemplate").ParseFiles("./static/taskList.html"))

func main() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	s := Server{
		db: db,
	}

	router := http.NewServeMux()
	router.HandleFunc("/", s.viewLists)
	router.HandleFunc("/list", s.viewList)

	port := "8080"
	log.Printf("server start at port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
