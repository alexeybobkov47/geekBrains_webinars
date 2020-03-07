package controllers

import (
	"log"
	"not-for-work/GeekBrainsWebinars/current-lessons/5/user-example/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
	Ormer orm.Ormer
}

func (c *UserController) Get() {
	users := make([]models.User, 0, 1)
	_, err := c.Ormer.QueryTable("user").All(&users)
	if err != nil {
		log.Print(err)
		return
	}

	log.Println(users)

	c.Data["Hello"] = "Hello, World"
	c.Data["Users"] = users

	c.TplName = "users.tpl"
}
