package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"not-for-work/GeekBrainsWebinars/current-lessons/8/posts-logging/models"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"

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

func (e Explorer) getList(listName string) (models.List, error) {
	c := e.Db.Database(e.DbName).Collection("lists")

	filter := bson.D{{Key: "name", Value: listName}}

	res := c.FindOne(context.Background(), filter)

	list := new(models.List)
	if err := res.Decode(list); err != nil {
		return models.List{}, err
	}
	return *list, nil
}

type taskReq struct {
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

/*
	curl -vX POST -H "Content-Type: application/json"  -d'{"text":"created task","complete":true}' http://localhost:8090/list?list_name=NewTest
*/

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

func (e Explorer) createTask(listName string, task models.Task) error {
	filter := bson.D{{Key: "name", Value: listName}}

	update := bson.D{{"$push", bson.D{{"list", task}}}}

	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.UpdateOne(context.Background(), filter, update)

	return err
}

/*
 curl -vX PUT -H "Content-Type: application/json"  -d'{"text":"NewTest_put","complete":false}' "http://localhost:8090/list?list_name=NewTest&id=0"
*/

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

func (e Explorer) updateTask(taskId, listName, newText string, complete bool) error {
	update := bson.D{}
	if len(newText) != 0 {
		update = append(update, bson.E{Key: fmt.Sprintf("list.%v.text", taskId), Value: newText})
	}
	update = append(update, bson.E{Key: fmt.Sprintf("list.%v.complete", taskId), Value: complete})

	update = bson.D{{Key: "$set", Value: update}}

	filter := bson.D{{Key: "name", Value: listName}}

	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.UpdateOne(context.Background(), filter, update)

	return err
}

/*
 curl -vX DELETE  "http://localhost:8090/list?id=10&list_name=NewTest"
*/

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

func (e Explorer) deleteTask(id, listName string) error {
	filter := bson.D{{Key: "name", Value: listName}}

	update := bson.D{
		{
			"$unset",
			bson.D{
				{Key: fmt.Sprintf("list.%v", id), Value: nil},
			},
		},
	}

	pull := bson.D{{Key: "$pull", Value: bson.D{{"list", nil}}}}

	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	_, err = c.UpdateOne(context.Background(), filter, pull)

	return err
}
