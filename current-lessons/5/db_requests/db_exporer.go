package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DSN = "root:1234@tcp(localhost:3306)/task_list_app?charset=utf8"

type Explorer struct {
	db *sql.DB
}

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

func main() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	e := Explorer{db: db}
	taskList := TaskList{
		ID:          3,
		Name:        "Test",
		Description: "New",
		List: []Task{{
			ID:       "10",
			ListID:   100,
			Text:     "New Text",
			Complete: false,
		},
		},
	}

	if err := e.CreateList(taskList); err != nil {
		log.Print(err)
		return
	}

	if err := e.CreateTask(taskList.ID, taskList.List[0]); err != nil {
		log.Print(err)
		return
	}

	if err := e.UpdateTask(5, "new text", false); err != nil {
		log.Print(err)
		return
	}
}

// CreateList — создание листа задач
func (e Explorer) CreateList(list TaskList) error {
	_, err := e.db.Exec("insert into task_list_app.lists (id, name, description) values (?,?,?)",
		list.ID, list.Name, list.Description)
	return err
}

// CreateTask — создание задачи
func (e Explorer) CreateTask(listId int, task Task) error {
	_, err := e.db.Exec("insert into task_list_app.tasks (list_id, text, complete) values (?, ?, ?)",
		listId, task.Text, task.Complete)

	return err
}

// UpdateTask - обновление задачи
func (e Explorer) UpdateTask(id int, text string, complete bool) error {
	if len(text) != 0 {
		_, err := e.db.Exec("UPDATE `task_list_app`.`tasks` SET `text` = ? WHERE (`id` = ?)", text, id)
		if err != nil {
			return err
		}
	}

	_, err := e.db.Exec("UPDATE `task_list_app`.`tasks` SET `complete` = ? WHERE (`id` = ?)", complete, id)

	return err
}
