package userCon

import (
	"../../common/http"
	"../../common/msg"
	"../../common/response"
	"../../config"
	"../../model/commonStruct"
	"../../model/userModel"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/segmentio/encoding/json"
	"strings"
)

type User struct{}

func (a *User) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("POST", "/register", "RegisterNewUser")     //注册新用户
	h.Handle("POST", "/login/", "Login")                 //普通登录
	h.Handle("GET", "/wx/check/login", "WxProgramCheck") //小程序登录验证
	//h.Handle("POST", "/wx/program/login", "WxProgramLogin") //小程序登录
}

func (a *User) WxProgramCheck(ctx iris.Context) iris.Map {
	var (
		err      error
		b        []byte
		secret   = config.Config.Get("wxApp.secret").(string)[1:]
		app_id   = config.Config.Get("wxApp.app_id").(string)[1:]
		WxApp    = new(commonStruct.WxApp)
		code     = ctx.URLParam("code")
		token    string
		res_user = commonStruct.User{}
	)
	if code == "" {
		return response.FailResponse(struct{}{}, errors.New(msg.Msg6))
	}
	URL := strings.Join([]string{"https://api.weixin.qq.com/sns/jscode2session?appid=",
		app_id, "&secret=", secret, "&js_code=", code, "&grant_type=authorization_code"}, "")

	if b, err = http.GetRequestBytes(URL); err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	fmt.Println("b===", string(b))
	if err = json.Unmarshal([]byte(b), &WxApp); err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	res_user, token, err = userModel.WxProgramLogin(WxApp.OpenID)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	return response.SuccessAndToken(res_user, "SUCCESS", token)
	//return response.SuccessAndToken(WxApp, "SUCCESS", "")
}

//func (a *User) WxProgramLogin(ctx iris.Context) iris.Map {
//	var (
//		err   error
//		token string
//		res_user  = commonStruct.User{}
//	)
//	type request struct {
//		Openid string `json:"openid"`
//	}
//	values := request{}
//	err = ctx.ReadJSON(&values)
//	if err != nil {
//		return response.FailResponse(struct{}{}, err)
//	}
//res_user,token, err = userModel.WxProgramLogin(values.Openid)
//if err != nil {
//	return response.FailResponse(struct{}{}, err)
//}
//return response.SuccessAndToken(res_user, "SUCCESS", token)
//}

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
	if values.PassWord == msg.Empty{
		return response.FailResponse(struct{}{}, errors.New(msg.Msg10))
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
	if values.PassWord == msg.Empty || values.UserName == msg.Empty{
		return response.FailResponse(struct{}{}, errors.New(msg.Msg10))
	}
	if len(values.PassWord) > 12 || len(values.PassWord) < 8{
		return response.FailResponse(struct{}{}, errors.New(msg.Msg1))
	}
	if len(values.UserName) > 10 || len(values.PassWord) < 3{
		return response.FailResponse(struct{}{}, errors.New(msg.Msg11))
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
