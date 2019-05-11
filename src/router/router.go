package router

import (
	"controller"
	"controller/sql"
	"controller/view"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Router() *iris.Application {
	app := iris.Default()
	//for path, con_troller := range Parties {
	//	mvc.New(app.Party(path)).Handle(con_troller)
	//}
	app.Handle("GET", "/", view.IndexHtml)
	mvc.New(app.Party("/user")).
		Handle(new(controller.Login))

	mvc.New(app.Party("/mysql")).
		Handle(new(sql.Query))
	return app
}
