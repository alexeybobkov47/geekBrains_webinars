package main

import (
	"database/sql"
	"log"
	"net/http"
)

type Server struct {
	db *sql.DB
}

func (server *Server) viewLists(w http.ResponseWriter, r *http.Request) {
	lists, err := getAllLists(server.db)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "alllists", lists); err != nil {
		log.Println(err)
	}
}

func (server *Server) viewList(w http.ResponseWriter, r *http.Request) {
	list, err := getList(server.db, r.URL.Query().Get("id"))
	if err != nil {
		log.Print(err)
		w.WriteHeader(404)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "list", list); err != nil {
		log.Println(err)
	}
}
