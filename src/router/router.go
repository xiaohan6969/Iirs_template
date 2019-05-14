package router

import (
	"controller/homepage"
	"github.com/kataras/iris"
	"middleware/corsServer"
)

func Router() *iris.Application {
	app := iris.Default()

	app.WrapRouter(corsServer.Cors().ServeHTTP) // 跨域请求

	app.Handle("GET", "/", homepage.IndexHtml) //首页

	return app
}
