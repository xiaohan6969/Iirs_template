package main

import (
	"./config"
	"./middleware/unusual"
	"./model"
	"./router"
	"github.com/kataras/iris"
)

var (
	port      = config.Config.Get("master.port").(string)
	SET_LEVEl = config.Config.Get("master.setLevel").(string)
)

func main() {
	//初始化数据库所有表
	model.Init()

	//初始化应用
	app := iris.New()

	// 同时写文件日志与控制台日志
	//f := log.NewLogFile()
	//defer log.DealErr(f)
	//app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	//app.Use(log.RequestLogger())

	//错误处理
	app.Use(unusual.IrisRavenMiddleware)

	//控制台信息
	app.Logger().SetLevel(SET_LEVEl)

	//初始化路由
	app = router.Router()

	//端口绑定
	error.Error(app.Run(iris.Addr(port)))
}
