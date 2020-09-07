package userCon

import (
	"errors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"../../common"
	"../../model/userModel"
)

type User struct{}

func (a *User) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("POST", "/register/new/user", "RegisterNewUser") //注册新用户
}

//注册
func (a *User) RegisterNewUser(ctx iris.Context) (iris.Map) {
	var (
		err error
		token string
	)
	type request struct {
		UserName   string `json:"user_name"`    //用户名称
		PassWord   string `json:"pass_word"`    //用户密码
	}
	values := request{}
	err = ctx.ReadJSON(&values)
	if err != nil {
		return common.FailResponse(struct {}{},err)
	}
	if len(values.PassWord) > 10{
		return common.FailResponse(struct {}{},errors.New(common.Msg1))
	}
	token,err = userModel.RegisterNewUserModel(values.UserName,values.PassWord)
	if err != nil {
		return common.FailResponse(struct {}{},err)
	}
	return common.SuccessAndToken(struct {}{},"",token)
}
