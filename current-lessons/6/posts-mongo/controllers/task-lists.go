package controllers

import (
	"context"

	"not-for-work/GeekBrainsWebinars/current-lessons/6/posts-mongo/models"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
	Explorer Explorer
}

func (c *ListController) Get() {
	lists, err := c.Explorer.getAllLists()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Data["Lists"] = lists
	c.TplName = "lists.tpl"
}

func (e Explorer) getAllLists() ([]models.List, error) {
	c := e.Db.Database(e.DbName).Collection("lists")

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	lists := make([]models.List, 0, 1)
	if err := cur.All(context.Background(), &lists); err != nil {
		return nil, err
	}

	return lists, nil
}

type postRequest struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

/*
	curl -vX POST -H "Content-Type: application/json"  -d'{"name":"NewTextMongo","desc":"newDesc"}' http://localhost:8090/lists
*/

func (c *ListController) Post() {
	resp := new(postRequest)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	list := models.List{
		Name:        resp.Name,
		Description: resp.Description,
	}

	if err := c.Explorer.createList(list); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}

func (e Explorer) createList(list models.List) error {
	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.InsertOne(context.Background(), list)

	return err
}

/*
	curl -vX PUT -H "Content-Type: application/json"  -d'{"name":"NewTest_put","desc":"newDesc_put"}' http://localhost:8090/lists?list_name=Test
*/

func (c *ListController) Put() {
	listName := c.Ctx.Request.URL.Query().Get("list_name")

	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty list_name"))
		return
	}

	resp := new(postRequest)
	if err := readAndUnmarshall(resp, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	list := models.List{
		Name:        resp.Name,
		Description: resp.Description,
	}

	if err := c.Explorer.updateList(listName, list.Name, list.Description); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}

func (e Explorer) updateList(listName, newName, newDesc string) error {
	filter := bson.D{{"name", listName}}
	update := bson.D{}

	if len(newName) != 0 {
		update = append(update, bson.E{Key: "name", Value: newName})
	}

	if len(newDesc) != 0 {
		update = append(update, bson.E{Key: "description", Value: newDesc})
	}

	update = bson.D{{"$set", update}}

	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.UpdateOne(context.Background(), filter, update)

	return err
}

/*
	curl -vX DELETE  http://localhost:8090/lists?list_name=NewTest_put
*/

func (c *ListController) Delete() {
	listName := c.Ctx.Request.URL.Query().Get("list_name")
	if len(listName) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty list_name"))
		return
	}

	err := c.Explorer.deleteList(listName)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	_, _ = c.Ctx.ResponseWriter.Write([]byte(`SUCCESS`))
}

func (e Explorer) deleteList(listName string) error {
	filter := bson.D{{"name", listName}}

	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.DeleteOne(context.Background(), filter)
	return err
}
