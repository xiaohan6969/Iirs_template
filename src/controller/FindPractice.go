package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"net/http"
	"server"
)

type Find struct {
}

func (a *Find) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/find", "Find")

}
func (a *Find) Find(ctx iris.Context) iris.Map {
	return iris.Map{
		"message": http.StatusCreated,
		"code":    server.Config.Get("Other.MyServerName"),
	}
}
