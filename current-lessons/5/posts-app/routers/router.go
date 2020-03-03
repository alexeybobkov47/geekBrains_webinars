package routers

import (
	"database/sql"
	"fmt"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/5/posts-app/controllers"
	"not-for-work/GeekBrainsWebinars/current-lessons/5/posts-app/models"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

const dsnBeeGo = "root:1234@tcp(localhost:3306)/beego?charset=utf8"
const dsnTasks = "root:1234@tcp(localhost:3306)/task_list_app?charset=utf8"

func init() {
	// Регистрация модели
	orm.RegisterModel(new(models.User))

	// Установка параметров подключения к БД
	_ = orm.RegisterDataBase("default", "mysql", dsnBeeGo, 30)
	// _ = orm.RegisterDataBase("default", "mysql", dsnTasks, 30)

	beeorm := orm.NewOrm()

	db, err := sql.Open("mysql", dsnTasks)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	user := models.User{Name: "slene"}

	// Объявление элемента в таблицу
	id, err := beeorm.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// Обновление данных
	user.Name = "astaxie"
	num, err := beeorm.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	beego.Router("/user", &controllers.UserController{
		Controller: beego.Controller{},
		Ormer:      beeorm,
	})

	beego.Router("/", &controllers.MainController{})

	beego.Router("/lists", &controllers.ListController{
		Controller: beego.Controller{},
		Db:         db,
	})

	beego.Router("/list", &controllers.TasksController{
		Controller: beego.Controller{},
		Db:         db,
	})
}
