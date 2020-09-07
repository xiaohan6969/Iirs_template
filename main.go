package main

import (
	"./config"
	"./log"
	"./model"
	"./router"
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/kataras/iris"
	"runtime/debug"
)

var (
	port      = config.Config.Get("master.port").(string)
	SET_LEVEl = config.Config.Get("master.setLevel").(string)
)

func main() {
	model.Init() //初始化数据库
	app := iris.New()
	r, CLOSE := log.ILogger()
	defer CLOSE()
	app.Use(r)
	app.Use(irisRavenMiddleware)
	app.Logger().SetLevel(SET_LEVEl)
	app = router.Router()
	error.Error(app.Run(iris.Addr(port)))
}

func irisRavenMiddleware(ctx iris.Context) {
	w, r := ctx.ResponseWriter(), ctx.Request()
	defer func() {
		if rval := recover(); rval != nil {
			debug.PrintStack()
			rvalStr := fmt.Sprint(rval)
			packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
			raven.Capture(packet, nil)
			w.WriteHeader(iris.StatusInternalServerError)
		}
	}()
	ctx.Next()
}
