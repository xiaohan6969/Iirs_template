package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"net/http"
)

type Login struct{}

func (a *Login) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/login", "Get")

}
func (a *Login) Get() iris.Map {
	return iris.Map{
		"code": http.StatusOK,
	}
}
