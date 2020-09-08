package unusual

import (
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/kataras/iris"
	"runtime/debug"
)

func IrisRavenMiddleware(ctx iris.Context) {
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
