package main

import (
	"os"

	_ "not-for-work/GeekBrainsWebinars/current-lessons/5/posts-app/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

/*
func init() {
	// Регистрация модели
	orm.RegisterModel(new(models.User))

	// Установка параметров подключения к БД
	_ = orm.RegisterDataBase("default", "mysql", dsnBeeGo, 30)
}


*/
func main() {
	/*
		o := orm.NewOrm()
		user := models.User{Name: "slene"}

		// Объявление элемента в таблицу
		id, err := o.Insert(&user)
		fmt.Printf("ID: %d, ERR: %v\n", id, err)

		// Обновление данных
		user.Name = "astaxie"
		num, err := o.Update(&user)
		fmt.Printf("NUM: %d, ERR: %v\n", num, err)
		// Чтение
		u := models.User{Id: user.Id}
		err = o.Read(&u)
		fmt.Printf("ERR: %v\n", err)

		/*
			// Удаление
			num, err = o.Delete(&u)
			fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	*/

	beego.Run("localhost:" + os.Getenv("httpport"))
}
