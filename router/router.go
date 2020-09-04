package router

import (
	"../controller/homepage"
	"../controller/testTemplate"
	"../middleware/corsServer"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Router() *iris.Application {
	app := iris.Default()

	app.WrapRouter(corsServer.Cors().ServeHTTP) // 跨域请求

	app.Handle("GET", "/", homepage.IndexHtml) //首页

	mvc.New(app.Party("/miniProgram")).
		Handle(new(testTemplate.SqlNature))

	return app
}
