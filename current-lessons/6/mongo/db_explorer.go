package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (e explorer) createList(list List) error {
	c := e.db.Database(dbName).Collection("lists")
	_, err := c.InsertOne(context.Background(), list)

	return err
}

func (e explorer) getLists(obj interface{}) error {
	c := e.db.Database(dbName).Collection("lists")

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	return cur.All(context.Background(), obj)
}

func (e explorer) getList(listName string, obj interface{}) error {
	c := e.db.Database(dbName).Collection("lists")

	filter := bson.D{{Key: "name", Value: listName}}

	res := c.FindOne(context.Background(), filter)

	return res.Decode(obj)
}

func (e explorer) updateList(listname string, newName string, newDesc string) error {
	filter := bson.D{{Key: "name", Value: listname}}

	update := bson.D{}

	if len(newName) != 0 {
		update = append(update, bson.E{Key: "name", Value: newName})
	}
	if len(newDesc) != 0 {
		update = append(update, bson.E{Key: "description", Value: newDesc})
	}
	update = bson.D{{Key: "$set", Value: update}}

	c := e.db.Database(dbName).Collection("lists")
	_, err := c.UpdateOne(context.Background(), filter, update)

	return err
}

func (e explorer) deleteList(listName string) error {
	c := e.db.Database(dbName).Collection("lists")

	filter := bson.D{{Key: "name", Value: listName}}

	_, err := c.DeleteOne(context.Background(), filter)
	return err
}

func (e explorer) deleteAll() error {
	c := e.db.Database(dbName).Collection("lists")

	filter := bson.D{}

	_, err := c.DeleteMany(context.Background(), filter)

	return err
}
