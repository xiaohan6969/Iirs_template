package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type First struct{}

func (a *First) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/first", "First")
}

func (a *First) First(ctx iris.Context) iris.Map {
	return iris.Map{
		"message": "First-connect-success",
	}
}
