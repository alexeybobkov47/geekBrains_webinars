package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/models"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/astaxie/beego"
)

type TasksController struct {
	beego.Controller
	Explorer Explorer
}

type Explorer struct {
	Db     *mongo.Client
	DbName string
}

// @Title GetTask
// @Description Вернет страницу со статистикой по задачам
// @Success 200 {html}
// @Failure 500 body is empty
// @router / [get]
func (c *TasksController) Get() {
	listName := c.Ctx.Request.URL.Query().Get("list_name")
	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(`empty list_name`))
		return
	}

	list, err := c.Explorer.getList(listName)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Data["List"] = list
	c.Data["Tasks"] = list.List
	c.TplName = "list.tpl"
}

type taskReq struct {
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

/*
	curl -vX POST -H "Content-Type: application/json"  -d'{"text":"created task","complete":true}' http://localhost:8090/list?list_name=NewTest
*/

// @Title CreateTask
// @Description Создвание нового листа
// @Param  name      path   string true      "Имя листа"
// @Param  desc      path   string true      "Описание листа"
// @Success 200 body is empty
// @Failure 500 body is empty
// @router / [post]
func (c *TasksController) Post() {
	listName := c.Ctx.Request.URL.Query().Get("list_name")
	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(`empty list_name`))
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

	if err := c.Explorer.createTask(listName, task); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
}

func readAndUnmarshall(resp interface{}, body io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, resp); err != nil {
		return err
	}

	return nil
}

/*
 curl -vX PUT -H "Content-Type: application/json"  -d'{"text":"NewTest_put","complete":false}' "http://localhost:8090/list?list_name=NewTest&id=0"
*/

// @Title UpdateTAsk
// @Description Обновление листа
// @Param  listname      path   string true      "Текущее название листа"
// @Param  name      path   string true      "Новое название"
// @Param  desc      path   string true      "Новое описание"
// @Success 200 body is empty
// @Failure 500 body is empty
// @router / [put]
func (c *TasksController) Put() {
	query := c.Ctx.Request.URL.Query()

	id := query.Get("id")
	if len(id) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(`empty id`))
		return
	}

	listName := query.Get("list_name")
	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(`empty list_name`))
		return
	}

	resp := new(taskReq)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	if err := c.Explorer.updateTask(id, listName, resp.Text, resp.Complete); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
}

/*
 curl -vX DELETE  "http://localhost:8090/list?id=10&list_name=NewTest"
*/

// @Title DeleteTask
// @Description Удаление листа
// @Param  listname      path   string true      "Название листа"
// @Success 200 body is empty
// @Failure 403 body is empty
// @router / [delete]
func (c *TasksController) Delete() {
	query := c.Ctx.Request.URL.Query()
	id := query.Get("id")
	if len(id) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	listName := query.Get("list_name")
	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty list_name"))
		return
	}

	if err := c.Explorer.deleteTask(id, listName); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
}
