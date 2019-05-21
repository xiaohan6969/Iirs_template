package main

import (
	"./irisTemplate/config"
	"./irisTemplate/router"
	"github.com/kataras/iris"
)

var (
	port = config.Config.Get("master.port").(string)
)

func main() {
	app := router.Router()
	error.Error(app.Run(iris.Addr(port)))
}
