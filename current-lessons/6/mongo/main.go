package main

import (
	"context"
	"flag"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// install:
// go get go.mongodb.org/mongo-driver/mongo

type queryMode int

const (
	createList queryMode = iota + 1
	readLists
	readList
	updateList
	deleteList
	deleteAll
)

type explorer struct {
	db *mongo.Client
}

func main() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(context.Background())

	e := explorer{db: db}

	taskList := List{
		ID:          111,
		Name:        "New",
		Description: "NewD",
		List: []Task{
			{
				Id:       12,
				Text:     "New Text",
				Complete: false,
			},
		},
	}

	lists := make([]List, 0, 1)
	list := new(List)

	fl := flag.Int("mode", int(readList), "")
	flag.Parse()
	mode := queryMode(*fl)

	if err := e.doAction(mode, taskList, list, &lists); err != nil {
		log.Fatal(err)
	}

	if mode == readList {
		log.Printf("%+v", list)
	} else if mode == readLists {
		log.Printf("%+v", lists)
	}
}

func (e explorer) doAction(mode queryMode, taskList List, list, lists interface{}) error {
	var err error

	switch mode {
	case createList:
		err = e.createList(taskList)
	case readLists:
		err = e.getLists(lists)
	case readList:
		err = e.getList("mail.ru", list)
	case updateList:
		err = e.updateList("New", "NewTestName", "NewDescr")
	case deleteList:
		err = e.deleteList("NewTestName")
	case deleteAll:
		err = e.deleteAll()
	default:
		return errors.Errorf("unknown mode: %v", mode)
	}

	return err
}
