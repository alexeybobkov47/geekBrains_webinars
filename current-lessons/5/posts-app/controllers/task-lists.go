package controllers

import (
	"database/sql"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/5/posts-app/models"

	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
	Db *sql.DB
}

func (c *ListController) Get() {
	lists, err := getAllLists(c.Db)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.Data["Lists"] = lists
	c.TplName = "lists.tpl"
}

func getAllLists(db *sql.DB) ([]models.Lists, error) {
	rows, err := db.Query("select * from task_list_app.lists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]models.Lists, 0, 1)
	for rows.Next() {
		list := models.Lists{}

		if err := rows.Scan(&list.Id, &list.Name, &list.Description); err != nil {
			log.Println(err)
			continue
		}

		res = append(res, list)
	}

	return res, nil
}

type postRequest struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

/*
	curl -vX POST -H "Content-Type: application/json"  -d'{"name":"NewText",
"desc":"newdesc"}' http://localhost:8090/lists
*/

func (c *ListController) Post() {
	resp := new(postRequest)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	list := models.Lists{
		Name:        resp.Name,
		Description: resp.Description,
	}

	if err := createList(c.Db, list); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS\n`))
}

func createList(db *sql.DB, list models.Lists) error {
	_, err := db.Exec("insert into task_list_app.lists (name, description) values (?,?)", list.Name, list.Description)

	return err
}

/*
	curl -vX PUT -H "Content-Type: application/json"  -d'{"name":"NewText","desc":"newdesc_put"}' http://localhost:8090/lists?id=131
*/
func (c *ListController) Put() {
	id := c.Ctx.Request.URL.Query().Get("id")

	if len(id) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	resp := new(postRequest)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	list := models.Lists{
		Name:        resp.Name,
		Description: resp.Description,
	}

	if err := updateList(c.Db, id, list.Name, list.Description); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}

func updateList(db *sql.DB, id, name, description string) error {
	if len(name) == 0 && len(description) == 0 {
		return nil
	}

	_, err := db.Exec("UPDATE `task_list_app`.`lists` SET `name`=?, `description`=? WHERE (`id` = ?)",
		name, description, id)

	return err
}

/*
	curl -vX DELETE  http://localhost:8090/lists?id=133
*/

func (c *ListController) Delete() {
	id := c.Ctx.Request.URL.Query().Get("id")

	err := deleteList(c.Db, id)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}

func deleteList(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM task_list_app.lists WHERE `id`=?", id)
	if err != nil {
		return err
	}

	return nil
}
