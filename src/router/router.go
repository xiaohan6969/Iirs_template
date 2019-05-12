package router

import (
	"controller/homepage"
	"controller/login"
	"controller/sqlHandle"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Router() *iris.Application {
	app := iris.Default()
	//for path, con_troller := range Parties {
	//	mvc.New(app.Party(path)).Handle(con_troller)
	//}
	app.Handle("GET", "/", homepage.IndexHtml)
	mvc.New(app.Party("/user")).
		Handle(new(login.Login))

	mvc.New(app.Party("/mysql")).
		Handle(new(sqlHandle.Query))

	return app
}
