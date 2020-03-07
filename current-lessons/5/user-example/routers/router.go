package routers

import (
	"fmt"
	"log"

	"not-for-work/GeekBrainsWebinars/current-lessons/5/user-example/controllers"
	"not-for-work/GeekBrainsWebinars/current-lessons/5/user-example/models"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

const dsnBeeGo = "root:1234@tcp(localhost:3306)/beego?charset=utf8"

func init() {
	log.Println("here")

	// Регистрация модели
	orm.RegisterModel(new(models.User))

	_ = orm.RegisterDataBase("default", "mysql", dsnBeeGo)

	beeorm := orm.NewOrm()

	user := models.User{Name: "selena"}

	// Объявление элемента в таблицу
	id, err := beeorm.Insert(&user)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("ID: %d\n", id)

	// Обновление данных
	user.Name = "astaxie"
	num, err := beeorm.Update(&user)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("NUM: %d\n", num)

	u := models.User{Id: user.Id}
	if err := beeorm.Read(&u); err != nil {
		log.Println(err)
		return
	}
	log.Print(u.Name)

	/*
		num, err = beeorm.Delete(&u)
		if err != nil {
			log.Print(err)
			return
		}

	*/

	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{
		Controller: beego.Controller{},
		Ormer:      beeorm,
	})
}
