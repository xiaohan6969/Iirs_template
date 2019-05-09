package router

import (
	"controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
)

func Router() *iris.Application {
	app := iris.Default()
	preSetting(app)
	//for path, con_troller := range Parties {
	//	mvc.New(app.Party(path)).Handle(con_troller)
	//}
	//mvc.New(app.Party("/user")).Handle(new(controllers.UserController))
	mvc.New(app.Party("/user")).
		Handle(new(controller.Login)).
		Handle(new(controller.First))

	mvc.New(app.Party("/second")).
		Handle(new(controller.Secoud))

	mvc.New(app.Party("/config")).
		Handle(new(controller.Find))
	return app
}

func preSetting(app *iris.Application) {
	// 定义错误显示级别
	app.Logger().SetLevel("debug")
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(
		customLogger,
		//recover2.New(),
	)
}
