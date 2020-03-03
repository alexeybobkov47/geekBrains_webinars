package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/5/posts-app/models"

	"github.com/astaxie/beego"
)

type TasksController struct {
	beego.Controller
	Db *sql.DB
}

func (c *TasksController) Get() {
	id := c.Ctx.Request.URL.Query().Get("id")
	if len(id) == 0 {
		log.Print("empty id")
		return
	}

	list, err := getList(c.Db, id)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		return
	}

	c.Data["List"] = list
	c.Data["Tasks"] = list.List
	c.TplName = "list.tpl"
}

// getList — получение списка по id
func getList(db *sql.DB, id string) (models.Lists, error) {
	list := models.Lists{}

	row := db.QueryRow(fmt.Sprintf("select * from task_list_app.lists where lists.id = %v", id))
	err := row.Scan(&list.Id, &list.Name, &list.Description)
	if err != nil {
		return list, err
	}

	rows, err := db.Query(fmt.Sprintf("select * from task_list_app.tasks WHERE tasks.list_id = %v", id))
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		task := models.Task{}

		err := rows.Scan(&task.Id, new(int), &task.Text, &task.Complete)
		if err != nil {
			log.Println(err)
			continue
		}

		list.List = append(list.List, task)
	}

	return list, nil
}

type taskReq struct {
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

/*
	curl -vX POST -H "Content-Type: application/json"  -d'{"text":"created task", "complete":true}' http://localhost:8090/list?id=3
*/

func (c *TasksController) Post() {
	id := c.Ctx.Request.URL.Query().Get("id")
	if len(id) == 0 {
		log.Print("empty id")
		return
	}

	resp := new(taskReq)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	task := models.Task{
		Text:     resp.Text,
		Complete: resp.Complete,
	}

	if err := createTask(c.Db, id, task); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
}

func readAndUnmarshall(resp interface{}, body io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Print("empty id")
		return err
	}

	if err := json.Unmarshal(bytes, resp); err != nil {
		return err
	}

	return nil
}

func createTask(db *sql.DB, id string, task models.Task) error {
	_, err := db.Exec("insert into task_list_app.tasks (list_id, text, complete) values (?, ?, ?)",
		id, task.Text, task.Complete)

	return err
}

/*
 curl -vX PUT -H "Content-Type: application/json"  -d'{"text":"NewTask", "complete":true}' http://localhost:8090/list?id=7
*/
func (c *TasksController) Put() {
	id := c.Ctx.Request.URL.Query().Get("id")
	if len(id) == 0 {
		log.Print("empty id")
		return
	}

	resp := new(taskReq)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	if len(resp.Text) == 0 {
		return
	}

	if err := updateTask(c.Db, id, resp.Text, resp.Complete); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
}

func updateTask(db *sql.DB, id string, text string, complete bool) error {
	if len(text) != 0 {
		_, err := db.Exec("UPDATE `task_list_app`.`tasks` SET `text` = ? WHERE (`id` = ?)", text, id)
		if err != nil {
			return err
		}
	}

	_, err := db.Exec("UPDATE `task_list_app`.`tasks` SET `complete` = ? WHERE (`id` = ?)", complete, id)

	return err
}

/*
 curl -vX DELETE  http://localhost:8090/list?id=8
*/

func (c *TasksController) Delete() {
	id := c.Ctx.Request.URL.Query().Get("id")
	if len(id) == 0 {
		log.Print("empty id")
		return
	}

	err := deleteTask(c.Db, id)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
}

func deleteTask(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM task_list_app.tasks WHERE `id`=?;", id)

	return err
}
