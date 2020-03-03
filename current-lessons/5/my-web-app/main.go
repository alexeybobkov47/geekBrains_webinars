package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Описние модели
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// Регистрация модели
	orm.RegisterModel(new(User))

	// Установка параметров подключения к БД
	orm.RegisterDataBase("default", "MySQL", "root:root@/my_db?charset=utf8", 30)
}

func main() {
	beego.Run("localhost")
}
