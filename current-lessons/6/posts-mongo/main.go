package main

import (
	_ "not-for-work/GeekBrainsWebinars/current-lessons/6/posts-mongo/routers"
	"os"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run("localhost:" + os.Getenv("httpport"))
}
