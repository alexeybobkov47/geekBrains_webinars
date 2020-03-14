package controllers

import (
	"context"
	"fmt"

	"not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/models"

	"go.mongodb.org/mongo-driver/bson"
)

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

func (e Explorer) createList(list models.List) error {
	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.InsertOne(context.Background(), list)

	return err
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

func (e Explorer) deleteList(listName string) error {
	filter := bson.D{{"name", listName}}

	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.DeleteOne(context.Background(), filter)
	return err
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

func (e Explorer) createTask(listName string, task models.Task) error {
	filter := bson.D{{Key: "name", Value: listName}}

	update := bson.D{{"$push", bson.D{{"list", task}}}}

	c := e.Db.Database(e.DbName).Collection("lists")
	_, err := c.UpdateOne(context.Background(), filter, update)

	return err
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
