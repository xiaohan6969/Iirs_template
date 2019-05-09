package main

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"router"
)

func main() {
	app := router.Router()
	golog.Info() //暂时不知道干啥的
	fmt.Println(app.Run(iris.Addr(":8080")))
}
