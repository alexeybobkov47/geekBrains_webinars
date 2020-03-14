// @APIVersion 1.0.0
// @Title Task List Documentation
// @Description Документация по работе с сервисом
// @Contact abc@nsnow.xyz
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"context"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers"

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

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/static",
			beego.NSInclude(
				&controllers.StatisticController{
					Explorer: controllers.Explorer{
						Db:     db,
						DbName: dbName,
					},
				},
			),
		),
		beego.NSNamespace("/lists",
			beego.NSInclude(
				&controllers.ListController{
					Explorer: controllers.Explorer{
						Db:     db,
						DbName: dbName,
					},
				},
			),
		),
		beego.NSNamespace("/list",
			beego.NSInclude(
				&controllers.TasksController{
					Explorer: controllers.Explorer{
						Db:     db,
						DbName: dbName,
					},
				},
			),
		),
	)

	beego.AddNamespace(ns)

	/*
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

	*/
}
