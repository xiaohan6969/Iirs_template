package userCon

import (
	"../../common/msg"
	"../../common/response"
	"../../model/userModel"
	"errors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type User struct{}

func (a *User) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("POST", "/register", "RegisterNewUser") //注册新用户
	h.Handle("POST", "/login", "Login")              //登录
}

//登录
func (a *User) Login(ctx iris.Context) iris.Map {
	var (
		err   error
		token string
	)
	type request struct {
		UserName string `json:"user_name"` //用户名称
		PassWord string `json:"pass_word"` //用户密码
	}
	values := request{}
	err = ctx.ReadJSON(&values)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	token, err = userModel.LoginModel(values.UserName, values.PassWord)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	return response.SuccessAndToken(struct{}{}, "SUCCESS", token)
}

//注册
func (a *User) RegisterNewUser(ctx iris.Context) iris.Map {
	var (
		err error
	)
	type request struct {
		UserName string `json:"user_name"` //用户名称
		PassWord string `json:"pass_word"` //用户密码
	}
	values := request{}
	err = ctx.ReadJSON(&values)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	if len(values.PassWord) > 10 {
		return response.FailResponse(struct{}{}, errors.New(msg.Msg1))
	}
	err = userModel.RegisterNewUserModel(values.UserName, values.PassWord)
	if err != nil {
		if err.Error()[0:10] == msg.Msg2 {
			return response.FailResponse(struct{}{}, errors.New(msg.Msg3))
		}
		return response.FailResponse(struct{}{}, err)
	}
	return response.SuccessAndToken(struct{}{}, "SUCCESS", "")
}
