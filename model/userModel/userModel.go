package userModel

import (
	"../../common"
	"../../middleware/jwt"
	"../../server/mysqlServer"
	"../commonStruct"
)

func RegisterNewUserModel(user_name,pass_word string) (string,error) {
	var (
		db = mysqlServer.JzGorm
		token string
	)
	token,_ = jwt.CreateToken(&jwt.Claims{UserName: user_name,})
	user := &commonStruct.User{
		UserName:user_name,
		PassWord:pass_word,
		UserAge:"",
		UserSex:"",
		CreateTime:common.TimeNowToStr(),
		Token:token,
	}
	return token,db.Create(user).Error
}
