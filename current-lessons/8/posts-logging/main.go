package main

import (
	"log"
	"os"

	_ "not-for-work/GeekBrainsWebinars/current-lessons/8/posts-logging/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

/*
	beego.Debug("Сообщение при дебаге приложения")
	beego.Info("Запись информации")
	beego.Notice("Уведомление")
	beego.Warn("Требует внимания")
	beego.Error("Ошибка")
	beego.Critical("Критическая ошибка")
	beego.Alert("Тревога")
	beego.Emergency("Чрезвычайная ситуация")
*/

func main() {
	if err := logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`); err != nil {
		log.Print(err)
	}

	beego.Notice("Уведомление")
	beego.Warn("Требует внимания")
	beego.Error("Ошибка")
	beego.Critical("Критическая ошибка")
	beego.Alert("Тревога")
	beego.Emergency("Чрезвычайная ситуация")
	beego.Info("Запись информации")
	beego.Debug("Сообщение при дебаге приложения")

	beego.Run("localhost:" + os.Getenv("httpport"))
}
