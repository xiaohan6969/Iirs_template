package userModel

import (
	"../../common/public"
	"../../common/msg"
	"../../middleware/jwt"
	"../../server/mysqlServer"
	"../commonStruct"
	"errors"
)

func RegisterNewUserModel(user_name, pass_word string) ( error) {
	var (
		db    = mysqlServer.JzGorm
	)
	user := &commonStruct.User{
		UserName:   user_name,
		PassWord:   pass_word,
		UserAge:    "",
		UserSex:    "",
		CreateTime: public.TimeNowToStr(),
		Token:      "",
	}
	return db.Create(user).Error
}

func LoginModel(user_name, pass_word string) (string, error){
	var (
		db    = mysqlServer.JzGorm
		token string
		err error
	)
	token, _ = jwt.CreateToken(&jwt.Claims{UserName: user_name})
	var count int
	err = db.Model(&commonStruct.User{}).
		Where(&commonStruct.User{UserName: user_name, PassWord: pass_word,}).
		Count(&count).Error
	if err != nil {
		return "", err
	}
	if count != 1{
		return "",errors.New(msg.Msg5)
	}
	return token,nil
}
