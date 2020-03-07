package routers

import (
	"database/sql"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/5/posts-app/controllers"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsnTasks = "root:1234@tcp(localhost:3306)/task_list_app?charset=utf8"
)

func init() {
	db, err := sql.Open("mysql", dsnTasks)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("db connected")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Router("/lists", &controllers.ListController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/list", &controllers.TasksController{
		Controller: beego.Controller{},
		Db:         db,
	})
}
