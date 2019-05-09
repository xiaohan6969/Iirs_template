package main

import (
	"github.com/kataras/iris"
	"router"
)

func main() {
	app := router.Router()
	error.Error(app.Run(iris.Addr(":8080")))
}
