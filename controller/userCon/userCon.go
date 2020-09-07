package userCon

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type User struct{}

func (a *User) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("POST", "/register/new/user", "RegisterNewUser") //注册新用户
}

//注册
func (a *User) RegisterNewUser(ctx iris.Context) (res int) {
	var ()
	res = 1
	return
}
