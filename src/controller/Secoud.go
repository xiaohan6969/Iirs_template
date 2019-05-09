package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"net/http"
)

type Secoud struct{}

func (a *Secoud) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/2", "Second")

}
func (a *Secoud) Second(ctx iris.Context) iris.Map {
	return iris.Map{
		"message": "The Second Success",
		"code":    http.StatusOK,
	}
}
