package router

import (
	"../controller/homepage"
	"../controller/indexCon"
	"../controller/userCon"
	"../middleware/corsServer"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Router() *iris.Application {
	app := iris.Default()
	app.WrapRouter(corsServer.Cors().ServeHTTP) // 跨域请求

	app.Handle("GET", "/", homepage.IndexHtml) //首页

	mvc.New(app.Party("/miniProgram")).
		Handle(new(indexCon.SqlNature))

	mvc.New(app.Party("/mini/user")).
		Handle(new(userCon.User))
	return app
}
