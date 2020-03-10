package routers

import (
	"context"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/6/posts-mongo/controllers"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "task_list_app"

func init() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongo-db connected")

	if err = db.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	beego.Router("/", &controllers.StatisticController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	})

	beego.Router("/lists", &controllers.ListController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	})

	beego.Router("/list", &controllers.TasksController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	})
}
