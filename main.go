package main

import (
	"./config"
	"./model"
	"./router"
	"github.com/kataras/iris"
)

var (
	port = config.Config.Get("master.port").(string)
)

func main() {
	model.Init() //初始化数据库
	app := router.Router()
	error.Error(app.Run(iris.Addr(port)))
}
