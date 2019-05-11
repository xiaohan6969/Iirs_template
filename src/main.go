package main

import (
	"./config"
	"github.com/kataras/iris"
	"router"
)

var (
	port = config.Config.Get("master.port").(string)
)

func main() {
	app := router.Router()
	error.Error(app.Run(iris.Addr(":" + port)))
}
