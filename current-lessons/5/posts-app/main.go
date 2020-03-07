package main

import (
	"os"

	_ "not-for-work/GeekBrainsWebinars/current-lessons/5/posts-app/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run("localhost:" + os.Getenv("httpport"))
}
