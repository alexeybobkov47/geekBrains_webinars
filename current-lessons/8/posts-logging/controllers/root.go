package controllers

import (
	"context"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/8/posts-logging/models"

	"github.com/pkg/errors"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
)

type StatisticController struct {
	beego.Controller
	Explorer Explorer
}

func (c *StatisticController) Get() {
	lists := make([]models.List, 0, 1)

	lists, err := c.Explorer.getLists()
	if err != nil {
		log.Print(err)
		return
	}

	var (
		countLists int
		countTasks int
		countOpen  int
		countClose int
	)

	for _, list := range lists {
		countLists++
		for _, task := range list.List {
			countTasks++
			if task.Complete {
				countClose++
				continue
			}
			countOpen++
		}
	}

	c.Data["Lists"] = countLists
	c.Data["Tasks"] = countTasks
	c.Data["Open"] = countOpen
	c.Data["Close"] = countClose

	c.TplName = "index.tpl"
}

func (e Explorer) getLists() ([]models.List, error) {
	c := e.Db.Database(e.DbName).Collection("lists")

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "Find")
	}

	lists := make([]models.List, 0, 1)
	if err := cur.All(context.Background(), &lists); err != nil {
		return nil, errors.Wrap(err, "All")
	}

	return lists, nil
}
